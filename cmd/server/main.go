package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"os"

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
		result := map[string]interface{}{}

		{
			res, err := c.Lemma.Query().
				Where(lemma.WrittenFormHasPrefix(query)).
				WithSenses().
				Limit(10).
				All(ctx)
			if err != nil {
				w.WriteHeader(500)
				fmt.Fprint(w, err.Error())
				return
			}
			for _, l := range res {
				result[l.WrittenForm] = l.Edges.Senses[0].QuerySynset().WithDefinitions().AllX(ctx)
			}
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
