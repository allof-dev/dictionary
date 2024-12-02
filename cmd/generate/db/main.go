package main

import (
	"context"
	"encoding/xml"
	"fmt"
	"os"

	"github.com/allof-dev/dictionary/ent"
	"github.com/allof-dev/dictionary/pkg/utils"
	"github.com/allof-dev/dictionary/pkg/wordnet"
	_ "github.com/mattn/go-sqlite3"
)

var (
	XMLFileName = os.Args[1]
	DBName      = os.Args[2]
)

func main() {
	r := wordnet.LexicalResource{}

	f, err := os.Open(XMLFileName)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	if err := xml.NewDecoder(f).Decode(&r); err != nil {
		panic(err)
	}

	lexcion := r.Lexicon[0]
	ctx := context.Background()
	if err := run(ctx, lexcion); err != nil {
		panic(err)
	}

}

func run(ctx context.Context, lexicon wordnet.Lexicon) error {
	c, err := utils.NewDBClient(DBName)
	if err != nil {
		return fmt.Errorf("failed to open database: %w", err)
	}
	defer c.Close()

	if err := c.Schema.Create(ctx); err != nil {
		return fmt.Errorf("fail to auto migration: %w", err)
	}

	tx, _ := c.Tx(ctx)
	{
		fmt.Println("Creating lemmas...")
		for _, l := range lexicon.LexicalEntries {
			_, err := tx.Lemma.Create().
				SetID(l.ID).
				SetPartOfSpeech(l.Lemma.PartOfSpeech).
				SetWrittenForm(l.Lemma.WrittenForm).
				Save(ctx)
			if err != nil {
				return fmt.Errorf("failed to save lemma: %w", err)
			}
		}
	}
	{
		fmt.Println("Creating synsets...")
		for _, s := range lexicon.Synsets {
			definitions := make([]int, len(s.Definitions))
			for i, def := range s.Definitions {
				if s, err := tx.Definition.Create().SetText(def).Save(ctx); err != nil {
					return fmt.Errorf("failed to save definition: %w", err)
				} else {
					definitions[i] = s.ID
				}
			}

			req := tx.Synset.Create()
			req.SetID(s.ID)
			req.AddDefinitionIDs(definitions...)
			req.SetPartOfSpeech(s.PartOfSpeech)
			if _, err := req.Save(ctx); err != nil {
				return fmt.Errorf("failed to save synset: %w", err)
			}
		}
	}
	{
		fmt.Println("Creating senses...")
		for _, l := range lexicon.LexicalEntries {
			for _, s := range l.Senses {
				_, err := tx.Sense.Create().
					SetID(s.ID).
					SetLemmaID(l.ID).
					SetSynsetID(s.Synset).
					Save(ctx)
				if err != nil {
					return fmt.Errorf("failed to save sense: %w", err)
				}
			}
		}
	}
	{
		fmt.Println("Creating sense relations")
		for _, l := range lexicon.LexicalEntries {
			for _, s := range l.Senses {
				var reqs []*ent.SenseRelationCreate
				for _, rel := range s.SenseRelations {
					reqs = append(reqs,
						tx.SenseRelation.Create().
							SetRelType(rel.RelType).
							SetFromID(s.ID).
							SetToID(rel.Target),
					)
				}
				if _, err := tx.SenseRelation.CreateBulk(reqs...).Save(ctx); err != nil {
					return fmt.Errorf("failed to create sense relations: %w", err)
				}
			}
		}
	}
	{
		fmt.Println("Creating synset relations")
		for _, s := range lexicon.Synsets {
			var reqs []*ent.SynsetRelationCreate
			for _, relation := range s.SynsetRelations {
				reqs = append(reqs, tx.SynsetRelation.Create().
					SetFromID(s.ID).
					SetToID(relation.Target).
					SetRelType(relation.RelType))
			}
			if _, err := tx.SynsetRelation.CreateBulk(reqs...).Save(ctx); err != nil {
				return fmt.Errorf("failed to create synset relations: %w", err)
			}
		}
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}
