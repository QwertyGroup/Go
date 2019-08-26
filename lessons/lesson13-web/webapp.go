package main

import (
	"encoding/xml"
	"io/ioutil"
	"net/http"
)

// Sitemapindex root node
type Sitemapindex struct {
	Locations []string `xml:"sitemap>loc"`
}

// News news
type News struct {
	Titles    []string `xml:"url>news>title"`
	Keywords  []string `xml:"url>news>Keywords"`
	Locations []string `xml:"url>loc"`
}

func main() {
	var s Sitemapindex
	var n News

	sitemapURL := "https://www.washingtonpost.com/news-sitemaps/index.xml"
	resp, _ := http.Get(sitemapURL)
	bytes, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	xml.Unmarshal(bytes, &s)

	for _, loc := range s.Locations {
		resp, _ := http.Get(loc)
		bytes, _ := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		xml.Unmarshal(bytes, &n)
	}
}
