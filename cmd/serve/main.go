package main

import (
	"context"
	"embed"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	create_receipt_v1 "github.com/sunb26/xat/handler/create_receipt"
	create_user_v1 "github.com/sunb26/xat/handler/create_user"
	get_receipt_v1 "github.com/sunb26/xat/handler/get_receipt"
)

//go:embed all:web all:web/_next
var content embed.FS

type middleware struct {
	handler http.Handler
	db      *sqlx.DB
}

func (i *middleware) ServeHTTP(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), "db", i.db)
		log.Printf("request header: %s %s", r.Method, r.URL)
		next.ServeHTTP(w, r.WithContext(ctx))
		log.Printf("response header: %#v", w.Header())
	})
}

func newMiddleware(handler http.Handler, db *sqlx.DB) *middleware {
	return &middleware{handler: handler, db: db}
}

func main() {
	content, err := fs.Sub(content, "web")
	if err != nil {
		log.Fatal(err)
	}

	db, err := sqlx.Connect("postgres", os.Getenv("DSN"))
	if err != nil {
		log.Fatalf("failed to open to database: %v", err)
	}

	// Use a combination of ServeMux and Router because Router does not handle FileServer well
	topMux := http.NewServeMux()
	prefixMux := mux.NewRouter()
	userMux := prefixMux.PathPrefix("/users/").Subrouter()

	mw := newMiddleware(userMux, db)
	userMux.Use(mw.ServeHTTP)

	userMux.HandleFunc("/v1/user", create_user_v1.CreateUser).Methods("PUT")
	userMux.HandleFunc("/v1/receipt", get_receipt_v1.GetReceipt).Methods("GET")
	userMux.HandleFunc("/v1/receipt", create_receipt_v1.CreateReceipt).Methods("PUT")

	topMux.Handle("/api/", http.StripPrefix("/api", prefixMux))
	topMux.Handle("/", http.FileServer(http.FS(content)))

	fmt.Println("Listening on 127.0.0.1:3000")
	log.Fatal(http.ListenAndServe(":3000", topMux))
}
