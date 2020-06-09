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
		indexer.Put(makeKey(item.Code), item)
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

func makeKey(code string) string {
	runes := []rune(code)

	if len(runes) <= 1 {
		return code
	}

	if len(runes) > 6 {
		runes = runes[0:6]
	}
	cut := len(runes)
	for i := cut - 1; i >= 1; {
		if runes[i] != '0' ||
			runes[i-1] != '0' {
			break
		}
		cut -= 2
		i -= 2
	}
	return string(runes[0:cut])
}

func AllDivisions() []string {
	names := []string{}
	for _, d := range items {
		names = append(names, d.Name)
	}
	return names
}
