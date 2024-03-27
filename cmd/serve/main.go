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

type injector struct {
	handler http.Handler
	db      *sqlx.DB
}

func (i *injector) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := context.WithValue(r.Context(), "db", i.db)
	i.handler.ServeHTTP(w, r.WithContext(ctx))
}

func newInjector(handler http.Handler, db *sqlx.DB) *injector {
	return &injector{handler: handler, db: db}
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

	mux := http.NewServeMux()
	mux.HandleFunc("/api/v1/user", create_user_v1.CreateUser)
	mux.Handle("/", http.FileServerFS(content))

	wrappedMux := newInjector(mux, db)

	fmt.Println("Listening on 127.0.0.1:3000")
	log.Fatal(http.ListenAndServe(":3000", wrappedMux))
}
