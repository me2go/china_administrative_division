package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/me2go/china_administrative_division/crawler"
)

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

	buf := bytes.NewBuffer([]byte{})
	buf.WriteString("package china_administrative_division")
	buf.WriteString("\n")
	buf.WriteString("var Divisions = ")
	buf.WriteString("`")
	buf.Write(data)
	buf.WriteString("`\n")

	ioutil.WriteFile("divisions.go", buf.Bytes(), os.FileMode(0644))
}
