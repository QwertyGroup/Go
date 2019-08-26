package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8000", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	html := `
<h1> Orsen </h1>
<p> Welcome </p>
`
	for i := 1; i <= 10; i++ {
		html += fmt.Sprintf("%d\n", i)
	}

	myArray := [...]string{"hello", " from", " Orsen"}
	for index, val := range myArray {
		html += fmt.Sprintf("%d -> %s\n", index, val)
	}

	mySlice := []string{"hello", " from", " Orsen"}
	for index, val := range mySlice {
		html += fmt.Sprintf("%d -> %s\n", index, val)
	}

	fmt.Fprint(w, html)
}
