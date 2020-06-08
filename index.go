package china_administrative_division

import (
	"encoding/json"

	"github.com/dghubble/trie"
	"github.com/me2go/china_administrative_division/crawler"
)

var indexer = trie.NewRuneTrie()
var divisionIndexer = trie.NewRuneTrie()

var items []crawler.Division

func init() {
	if err := json.Unmarshal([]byte(Divisions), &items); err != nil {
		panic(err)
	}
	for _, item := range items {
		indexer.Put(item.Code, item)
	}

	for _, item := range items {
		ds := []crawler.Division{}
		indexer.WalkPath(item.Code, func(key string, v interface{}) error {
			d, _ := v.(crawler.Division)
			ds = append(ds, d)
			return nil
		})
		if len(ds) > 1 {
			key := ""
			for _, d := range ds {
				key += d.Name
			}
			divisionIndexer.Put(key, ds)
		}
	}
}

func AllDivisions() []string {
	names := []string{}
	for _, d := range items {
		names = append(names, d.Name)
	}
	return names
}
