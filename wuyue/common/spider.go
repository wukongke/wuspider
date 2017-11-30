package common

import (
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"github.com/Tang-RoseChild/mahonia"
)

func Curl(url string) *goquery.Document {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("http err: ", err)
	}
	defer resp.Body.Close()

	dec := mahonia.NewDecoder("GB18030")
	rd := dec.NewReader(resp.Body)

	doc, err := goquery.NewDocumentFromReader(rd)
	if err != nil {

		fmt.Println("reader error %s ", err.Error())
	}
	return doc
}
