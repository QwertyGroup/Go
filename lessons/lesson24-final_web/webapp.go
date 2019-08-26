package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"
	"text/template"
)

var wg sync.WaitGroup

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

func newsRoutine(c chan News, i int, loc string) {
	defer wg.Done()
	var n News
	loc = strings.TrimSpace(loc)
	toXML(loc, &n)
	fmt.Printf("Got %d subnode\n", i+1)
	c <- n
}

func toXML(url string, v interface{}) {
	resp, _ := http.Get(url)
	bytes, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	xml.Unmarshal(bytes, v)
}

var s Sitemapindex

func loadSitemap() {
	sitemapURL := "https://www.washingtonpost.com/news-sitemaps/index.xml"
	toXML(sitemapURL, &s)
	fmt.Println("Got root node")
}

var cached = false

func newsAggHandler(w http.ResponseWriter, r *http.Request) {
	if !cached {
		loadSitemap()
		cached = true
	}

	queue := make(chan News, 100)     // just buffer more than needed
	for i, loc := range s.Locations { //[:5] {
		wg.Add(1)
		go newsRoutine(queue, i, loc)
	}

	wg.Wait()
	close(queue)

	newsMap := make(map[string]NewsMap)
	for elem := range queue {
		n := elem
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
