package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	sitemapURL := "https://www.washingtonpost.com/news-sitemaps/index.xml"
	resp, _ := http.Get(sitemapURL)
	bytes, _ := ioutil.ReadAll(resp.Body)
	stringBody := string(bytes)
	fmt.Println(stringBody)
	resp.Body.Close()
}
