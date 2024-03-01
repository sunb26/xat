package main

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
)

//go:embed all:web all:web/_next
var content embed.FS

func main() {
	content, err := fs.Sub(content, "web")
	if err != nil {
		log.Fatal(err)
	}
	http.Handle("/", http.FileServerFS(content))
	log.Fatal(http.ListenAndServe(":3000", nil))
}
