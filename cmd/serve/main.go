package main

import (
	"embed"
	"fmt"
	"io/fs"
	"log"
	"net/http"

	handler_v1 "github.com/sunb26/xat/handler"
)

//go:embed all:web all:web/_next
var content embed.FS

func main() {
	content, err := fs.Sub(content, "web")
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/api/v1/user", http.HandlerFunc(handler_v1.CreateUser))
	http.Handle("/", http.FileServerFS(content))

	fmt.Println("Listening on 127.0.0.1:3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
