package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/allof-dev/dictionary/ent"
	"github.com/allof-dev/dictionary/pkg/utils"
)

var (
	DBName     = os.Args[1]
	OUTPUTFile = os.Args[2]
)

func main() {
	if err := run(context.Background()); err != nil {
		panic(err)
	}
}

func run(ctx context.Context) error {
	f, err := os.Create(OUTPUTFile)
	if err != nil {
		return err
	}

	db, err := utils.NewDBClient(DBName)
	if err != nil {
		return err
	}

	for _, def := range db.Definition.Query().AllX(ctx) {
		var id string
		// get words for definition
		words := []string{}
		for _, synset := range def.
			QuerySynset().
			WithSense(func(sq *ent.SenseQuery) {
				sq.WithLemma()
			}).
			AllX(ctx) {
			id = synset.ID
			for _, sense := range synset.Edges.Sense {
				words = append(words, sense.Edges.Lemma.WrittenForm)
			}
		}

		req := generateReq(id, words, def.Text)
		if err := json.NewEncoder(f).Encode(req); err != nil {
			return err
		}
	}

	return nil
}

func generateReq(id string, words []string, definition string) OpenAIRequest {
	var req = OpenAIRequest{
		CustomID: id,
		Method:   "POST",
		URL:      "/v1/chat/completions",
		Body: ChatCompletionBody{
			Model:            "gpt-4o-mini",
			Temperature:      0,
			MaxTokens:        150,
			TopP:             1,
			FrequencyPenalty: 0,
			PresencePenalty:  0,
			Messages: []Messages{
				{
					Role: "system",
					Content: []Content{
						{Type: "text", Text: "주어진 Synset 단어 목록과 정의를 한국어로 번역합니다. 번역은 간결하고 직관적인 문체로 진행합니다."},
					},
				},
				{
					Role: "user",
					Content: []Content{
						{Type: "text", Text: fmt.Sprintf("%s\n%s", strings.Join(words, ", "), definition)},
					},
				},
			},
			ResponseFormat: ResponseFormat{
				Type: "json_schema",
				JSONSchema: JSONSchema{
					Name:   "synset_translation",
					Strict: true,
					Schema: Schema{
						Type: "object",
						Required: []string{
							"words",
							"definition",
						},
						AdditionalProperties: false,
						Properties: Properties{
							Words: Words{
								Type:        "array",
								Description: "An array of synset words in Korean.",
								Items: Items{
									Type:        "string",
									Description: "A Korean word in the synset.",
								},
							},
							Definition: Definition{
								Type:        "string",
								Description: "The definition of the synset in Korean.",
							},
						},
					},
				},
			},
		},
	}

	return req
}
