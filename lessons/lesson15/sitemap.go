package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// Sitemapindex root node
type Sitemapindex struct {
	Locations []string `xml:"sitemap>loc"`
}

// News news
type News struct {
	Titles    []string `xml:"url>news>title"`
	Keywords  []string `xml:"url>news>keywords"`
	Locations []string `xml:"url>loc"`
}

// NewsMap news map
type NewsMap struct {
	Keyword  string
	Location string
}

func main() {
	var s Sitemapindex
	var n News
	newsMap := make(map[string]NewsMap)

	sitemapURL := "https://www.washingtonpost.com/news-sitemaps/index.xml"
	resp, _ := http.Get(sitemapURL)
	bytes, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	xml.Unmarshal(bytes, &s)

	for _, loc := range s.Locations {
		loc = strings.TrimSpace(loc)
		resp, _ := http.Get(loc)
		bytes, _ := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		xml.Unmarshal(bytes, &n)

		for idx := range n.Titles {
			newsMap[n.Titles[idx]] = NewsMap{n.Keywords[idx], n.Locations[idx]}
		}

		for title, data := range newsMap {
			fmt.Println("\n\n\n", title)
			fmt.Println("\n", data.Keyword)
			fmt.Println("\n", data.Location)
		}
	}
}
