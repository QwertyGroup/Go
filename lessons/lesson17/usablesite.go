package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"text/template"
)

// NewsAggPage news page
type NewsAggPage struct {
	Title string
	News  map[string]NewsMap
}

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

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home Page | Welcome!")
}

func newsAggHandler(w http.ResponseWriter, r *http.Request) {
	var s Sitemapindex
	var n News
	newsMap := make(map[string]NewsMap)

	sitemapURL := "https://www.washingtonpost.com/news-sitemaps/index.xml"
	resp, _ := http.Get(sitemapURL)
	bytes, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	xml.Unmarshal(bytes, &s)
	fmt.Println("Got root node")

	for i, loc := range s.Locations[:5] {
		loc = strings.TrimSpace(loc)
		resp, _ := http.Get(loc)
		bytes, _ := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		xml.Unmarshal(bytes, &n)
		fmt.Printf("Got %d subnode\n", i+1)

		for idx := range n.Titles {
			newsMap[n.Titles[idx]] = NewsMap{n.Keywords[idx], n.Locations[idx]}
		}
	}

	p := NewsAggPage{Title: "CNTNT", News: newsMap}
	t, _ := template.ParseFiles("newstemplate.html")
	err := t.Execute(w, p)
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/agg/", newsAggHandler)
	http.ListenAndServe(":8001", nil)
}
