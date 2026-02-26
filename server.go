package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

type Page struct {
	Body  []byte
}

func loadPage(filename string) (*Page, error) {
	if filename == "" {
		filename = "index.html"
	}
	body, err := os.ReadFile("pages/" + filename)
	if err != nil {
		return nil, err
	}
	return &Page{Body: body}, nil
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	pagePath := r.URL.Path[1:]
	p, _ := loadPage(pagePath)
	fmt.Fprintf(w, "%s", p.Body)
}

func main() {
	http.HandleFunc("/", viewHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
