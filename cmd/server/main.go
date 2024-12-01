package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"entgo.io/ent/dialect/sql"
	"github.com/allof-dev/dictionary/ent"
	"github.com/allof-dev/dictionary/ent/lemma"
	_ "github.com/mattn/go-sqlite3"
)

var (
	DBName = os.Args[1]
)

func main() {
	ctx := context.Background()

	c, err := prepareDB()
	if err != nil {
		panic(err)
	}
	defer c.Close()

	if err := runServer(ctx, c); err != nil {
		panic(err)
	}

}

func prepareDB() (*ent.Client, error) {
	c, err := ent.Open("sqlite3", fmt.Sprintf("file:%s?_fk=1", DBName))
	if err != nil {
		return nil, fmt.Errorf("failed to open db: %w", err)
	}
	return c, nil
}

func runServer(ctx context.Context, c *ent.Client) error {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /search", func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query().Get("q")
		result, err := searchResult(ctx, c, query)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "error: %s", err.Error())
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		enc := json.NewEncoder(w)
		if err := enc.Encode(result); err != nil {
			slog.Error(err.Error())
		}
	})

	fmt.Println("http://127.0.0.1:8080/search")
	return http.ListenAndServe(":8080", mux)
}

type Lemma struct {
	ID           string   `json:"id"`
	WrittenFrom  string   `json:"written_form"`
	PartOfSpeech string   `json:"part_of_speech"`
	Synsets      []Synset `json:"synsets"`
}

type Synset struct {
	ID          string   `json:"id"`
	Words       []string `json:"word"`
	Definitions []string `json:"definitions"`
}

func searchResult(ctx context.Context, c *ent.Client, query string) (result []Lemma, err error) {
	tmp := []Lemma{}

	// find lemmas
	lemmas, err := c.Lemma.Query().
		Where(lemma.WrittenFormHasPrefix(query)).
		WithSenses(func(q *ent.SenseQuery) {
			q.WithSynset()
		}).
		Order(
			func(s *sql.Selector) {
				s.OrderExprFunc(func(b *sql.Builder) {
					b.WriteString(
						fmt.Sprintf("%s COLLATE NOCASE ASC", lemma.FieldWrittenForm),
					)
				})
			},
		).
		Limit(10).
		All(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to find lemmas: %w", err)
	}

	for _, lemma := range lemmas {
		l := Lemma{
			ID:           lemma.ID,
			WrittenFrom:  lemma.WrittenForm,
			PartOfSpeech: lemma.PartOfSpeech,
		}
		for _, sense := range lemma.Edges.Senses {
			synset := sense.Edges.Synset
			if synset == nil {
				continue
			}
			s := Synset{}

			// Fill ID
			s.ID = synset.ID

			// Fill Definitions
			definitions, err := synset.QueryDefinitions().All(ctx)
			if err != nil {
				return nil, err
			}
			for _, def := range definitions {
				s.Definitions = append(s.Definitions, def.Text)
			}

			// Fill Words
			senses, err := synset.
				QuerySense().
				WithLemma().
				All(ctx)
			if err != nil {
				return nil, err
			}
			for _, sense := range senses {
				s.Words = append(s.Words, sense.Edges.Lemma.WrittenForm)
			}

			l.Synsets = append(l.Synsets, s)
		}
		tmp = append(tmp, l)
	}

	return tmp, nil
}
