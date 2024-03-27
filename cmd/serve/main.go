package main

import (
	"context"
	"embed"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	create_user_v1 "github.com/sunb26/xat/handler/create_user"
)

//go:embed all:web all:web/_next
var content embed.FS

func injectDB(db *sqlx.DB, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), "db", db)

		next.ServeHTTP(w, r.WithContext(ctx))
	}
}

func main() {
	content, err := fs.Sub(content, "web")
	if err != nil {
		log.Fatal(err)
	}

	db, err := sqlx.Open("postgres", os.Getenv("DSN"))
	if err != nil {
		log.Fatalf("failed to open to database: %v", err)
	}

	http.HandleFunc("/api/v1/user", injectDB(db, create_user_v1.CreateUser))
	http.Handle("/", http.FileServerFS(content))

	fmt.Println("Listening on 127.0.0.1:3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
