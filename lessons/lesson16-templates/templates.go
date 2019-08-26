package main

import (
	"fmt"
	"net/http"
	"text/template"
)

// NewsAggPage news page
type NewsAggPage struct {
	Title string
	News  string
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello")
}

func newsAggHandler(w http.ResponseWriter, r *http.Request) {
	p := NewsAggPage{Title: "CNTNT", News: "some news"}
	t, _ := template.ParseFiles("basictemplate.html")
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
