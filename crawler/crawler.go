package crawler

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/kr/pretty"
	"github.com/parnurzeal/gorequest"
)

const pageUrl = "http://www.mca.gov.cn/article/sj/xzqh/2020/2020/202003301019.html"

type Division struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

func extractDivision(s *goquery.Selection) Division {
	if s.Size() < 2 {
		return Division{}
	}
	return Division{
		Code: strings.Trim(s.Eq(0).Text(), " "),
		Name: strings.Trim(s.Eq(1).Contents().Not("span").Text(), " "),
	}
}

func Crawler() ([]Division, error) {
	_, body, errs := gorequest.New().
		Get(pageUrl).
		End()
	if len(errs) != 0 {
		return nil, fmt.Errorf("get html page error: %v", pretty.Sprint(errs))
	}

	doc, err := goquery.NewDocumentFromReader(bytes.NewBufferString(body))
	if err != nil {
		return nil, err
	}

	divisions := []Division{}

	doc.Find("tr").Each(func(i int, s *goquery.Selection) {
		//省市一级
		child := s.Find("td[class=xl7030721]")
		if d := extractDivision(child); d.Code != "" {
			divisions = append(divisions, d)
		}
		//区县一级
		child = s.Find("td[class=xl7130721]")
		if d := extractDivision(child); d.Code != "" {
			divisions = append(divisions, d)
		}
	})

	return divisions, nil
}
