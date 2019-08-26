package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

// SitemapIndex root node
type SitemapIndex struct {
	Locations []Location `xml:"sitemap"`
}

// Location child node
type Location struct {
	Loc string `xml:"loc"`
}

func (l Location) String() string {
	return fmt.Sprintf(l.Loc)
}

func main() {
	sitemapURL := "https://www.washingtonpost.com/news-sitemaps/index.xml"
	resp, _ := http.Get(sitemapURL)
	bytes, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	var s SitemapIndex
	xml.Unmarshal(bytes, &s)
	// fmt.Println(s.Locations)
	for _, loc := range s.Locations {
		fmt.Printf("\n%s", loc)
	}
}
