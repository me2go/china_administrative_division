package china_administrative_division

import "github.com/me2go/china_administrative_division/crawler"

func Search(code string) ([]string, bool) {
	segs := []string{}
	indexer.WalkPath(code, func(k string, v interface{}) error {
		d, ok := v.(crawler.Division)
		if ok {
			segs = append(segs, d.Name)

		}
		return nil
	})
	if len(segs) > 0 {
		return segs, true
	}
	return nil, false
}
