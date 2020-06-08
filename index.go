package china_administrative_division

import (
	"encoding/json"

	"github.com/dghubble/trie"
	"github.com/me2go/china_administrative_division/crawler"
)

var indexer = trie.NewRuneTrie()

func init() {
	var items []crawler.Division
	if err := json.Unmarshal([]byte(Divisions), &items); err != nil {
		panic(err)
	}
	for _, item := range items {
		indexer.Put(item.Code, item)
	}
}
