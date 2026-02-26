package main

import (
	"github.com/sendakvmahr/go-web-spritesheets/spritesheets"
	"os"
	"log"
	"net/http"
	"fmt"
	"encoding/json"
)

type Page struct {
	Body  []byte
}

func loadPage(filename string) (*Page, error) {
	if filename == "" {
		filename = "index.html"
	} else {
		filename = filename + ".html"
	}
	body, err := os.ReadFile("pages/" + filename)
	if err != nil {
		return nil, err
	}
	return &Page{Body: body}, nil
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	pagePath := r.URL.Path[1:]
	p, err := loadPage(pagePath)
    if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "ehh custom 404 good enough")
        return
    }
	fmt.Fprintf(w, "%s", p.Body)
}

func apiHandler(w http.ResponseWriter, r *http.Request) {
	endpoint := r.URL.Path[len("/api/"):]
	fmt.Printf("%v\n", r.URL.Path)
	if endpoint == "fetchPixels" {
		handlePixelFetch(w, r)
	} else {
		fmt.Fprintf(w, "<h1>%s - you misspelled something</h1>", endpoint)
	}
}

func handlePixelFetch(w http.ResponseWriter, r *http.Request) {
	// pretty much mocked up to get everything in an okay points
	var ImageMap = map[string]spritesheets.SpriteSlice{}
	ImageMap["black_textbox_bg"] = spritesheets.SpriteSlice{
		Image: "textboxes",
		X : 14,
		Y: 17,
		Width: 2,
		Height: 2,
	}
	// plop the b64 image as is for now
	data := spritesheets.SpriteSheetResponse{
		Images : map[string]string{
			"textboxes" : "R0lGODdhEgATAHcAACH/C05FVFNDQVBFMi4wAwEAAAAh+QQJCgAAACwAAAAAEgATAIQAAAAAAAAHEAspVjIcKFYcTHZPQGlGnWZcaaZzic+RsKyW46ixqcG+vtGwxsHS4t7R0ebr9PD///8AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAFj2ATjWTUACjArGiUvi7KIAXCAPGLywgkQbYcTFaQGAsMYSo28wGTulZqVrsppVPWLio0eFGi0uiEMigcCgMgLB64zQqCIn0dHBYDg4MAIDgMSnYLB3lzcnRuAwCChGVnaTgDg4yKKV5qkXeaAzEJKDYIO5RbD54MElYoiVgPpaBrYmMopQABAlG4ngACASghADs=",
		},
		ImageDict: ImageMap,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func main() {
	http.HandleFunc("/", viewHandler)
	http.HandleFunc("/api/", apiHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
