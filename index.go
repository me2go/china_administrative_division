package china_administrative_division

import (
	"encoding/json"

	"github.com/dghubble/trie"
	"github.com/me2go/china_administrative_division/crawler"
)

var indexer = trie.NewRuneTrie()
var items []crawler.Division

func init() {
	if err := json.Unmarshal([]byte(Divisions), &items); err != nil {
		panic(err)
	}
	for _, item := range items {
		indexer.Put(item.Code, item)
	}
}

func AllDivisions() []string {
	names := []string{}
	for _, d := range items {
		names = append(names, d.Name)
	}
	return names
}
