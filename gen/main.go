package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/me2go/china-administrative-division/crawler"
)

var template = `package data
var Divisions = %s
`

func main() {
	divisions, err := crawler.Crawler()
	if err != nil {
		log.Println(err)
		return
	}
	data, err := json.Marshal(divisions)
	if err != nil {
		log.Println(err)
		return
	}

	content := fmt.Sprintf(template, string(data))
	ioutil.WriteFile("data/data.go", []byte(content), os.FileMode(0755))
}
