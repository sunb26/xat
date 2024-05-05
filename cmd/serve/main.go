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

	create_receipt_v1 "github.com/sunb26/xat/handler/create_receipt"
	create_scan_inference_v1 "github.com/sunb26/xat/handler/create_scan_inference"
	create_user_v1 "github.com/sunb26/xat/handler/create_user"
	get_receipt_v1 "github.com/sunb26/xat/handler/get_receipt"
	list_receipts_v1 "github.com/sunb26/xat/handler/list_receipts"
)

//go:embed all:web all:web/_next
var content embed.FS

type middleware struct {
	handler http.Handler
	db      *sqlx.DB
}

func (i *middleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := context.WithValue(r.Context(), "db", i.db)
	log.Printf("request header: %s %s", r.Method, r.URL)
	i.handler.ServeHTTP(w, r.WithContext(ctx))
	log.Printf("response header: %#v", w.Header())
func (i *middleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := context.WithValue(r.Context(), "db", i.db)
	log.Printf("request header: %s %s", r.Method, r.URL)
	i.handler.ServeHTTP(w, r.WithContext(ctx))
	log.Printf("response header: %#v", w.Header())
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

	topMux := http.NewServeMux()
	apiMux := http.NewServeMux()
	staticMux := http.NewServeMux()
	wrappedApiMux := newMiddleware(apiMux, db)

	apiMux.HandleFunc("PUT /v1/user", create_user_v1.CreateUser)
	apiMux.HandleFunc("GET /v1/receipt/{receiptId}", get_receipt_v1.GetReceipt)
	apiMux.HandleFunc("PUT /v1/receipt", create_receipt_v1.CreateReceipt)
	apiMux.HandleFunc("GET /v1/users/{userId}/receipts", list_receipts_v1.ListReceipts)
	apiMux.HandleFunc("PUT /v1/scan/inference", create_scan_inference_v1.CreateScanInference)
	staticMux.Handle("/", http.FileServer(http.FS(content)))

	topMux.Handle("/api/", http.StripPrefix("/api", wrappedApiMux))
	topMux.Handle("/", staticMux)
	topMux.Handle("/api/", http.StripPrefix("/api", wrappedApiMux))
	topMux.Handle("/", staticMux)

	fmt.Println("Listening on 127.0.0.1:3000")
	log.Fatal(http.ListenAndServe(":3000", topMux))
}
