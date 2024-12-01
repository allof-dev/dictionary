package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/allof-dev/dictionary/ent"
	_ "github.com/mattn/go-sqlite3"
	"github.com/ollama/ollama/api"
)

var (
	DBName     = os.Args[1]
	OUTPUTFile = os.Args[2]
)

func main() {
	ctx := context.Background()
	if err := run(ctx); err != nil {
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

type Result struct {
	ID   int `json:"id"`
	Word `json:""`
	Src  Word `json:"src"`
}

type Word struct {
	Meaning string   `json:"meaning"`
	Words   []string `json:"words"`
}

func (w *Word) print() {
	fmt.Printf("%s: %s\n", strings.Join(w.Words, ","), w.Meaning)
}

func run(ctx context.Context) error {
	f, err := os.Create(OUTPUTFile)
	if err != nil {
		return err
	}

	c, err := api.ClientFromEnvironment()
	if err != nil {
		return err
	}

	db, err := prepareDB()
	if err != nil {
		return err
	}

	defs := db.Definition.Query().
		AllX(ctx)

	for i, def := range defs {
		words := []string{}
		for _, synset := range def.QuerySynset().
			WithSense(func(sq *ent.SenseQuery) { sq.WithLemma() }).
			AllX(ctx) {
			for _, sense := range synset.Edges.Sense {
				words = append(words, sense.Edges.Lemma.WrittenForm)
			}
		}

		en := Word{
			Meaning: def.Text,
			Words:   words,
		}
		fmt.Printf("%06d/%d\n", i, len(defs))
		en.print()
		ko, err := generate(ctx, c, en)
		if err != nil {
			return err
		}
		ko.print()

		if err := json.NewEncoder(f).Encode(Result{
			ID:   def.ID,
			Word: ko,
			Src:  en,
		}); err != nil {
			return err
		}
	}

	return nil
}

func generate(ctx context.Context, c *api.Client, q Word) (Word, error) {

	b, err := json.Marshal(q)
	if err != nil {
		return Word{}, err
	}

	response := Word{}

	F := false
	if err := c.Generate(ctx, &api.GenerateRequest{
		Model:  "gemma2:9b",
		System: "주어진 JSON 의 value 를 한국어로 변환합니다.\n주어진 JSON 의 key 는 원문 그대로 유지합니다.\n고유명사는 원문을 유지하되, 이외에는 모두 한국어로 작성합니다.\n한국어 변환은 간결하고 직관적인 문체로 진행합니다.",
		Format: "json",
		Prompt: string(b),
		Stream: &F,
		Options: map[string]interface{}{
			"temperature": 0.0,
		},
	}, func(gr api.GenerateResponse) error {
		return json.Unmarshal([]byte(gr.Response), &response)
	}); err != nil {
		return Word{}, nil
	}
	return response, nil
}
