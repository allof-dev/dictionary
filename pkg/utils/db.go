package utils

import (
	"fmt"

	"github.com/allof-dev/dictionary/ent"
	_ "github.com/mattn/go-sqlite3"
)

func NewDBClient(DBFile string) (*ent.Client, error) {
	c, err := ent.Open("sqlite3", fmt.Sprintf("file:%s?_fk=1", DBFile))
	if err != nil {
		return nil, fmt.Errorf("failed to open db: %w", err)
	}
	return c, nil
}
