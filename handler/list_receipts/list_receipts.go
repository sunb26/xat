package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type receiptResponse struct {
	ProjectId uint64  `db:"project_id"`
	UserId    string  `db:"user_id"`
	ReceiptId uint64  `db:"receipt_id"`
	ScanId    *uint64 `db:"scan_id"`
	Subtotal  string  `db:"subtotal"`
	Tax       string  `db:"tax"`
	Gratuity  string  `db:"gratuity"`
	Date      string  `db:"receipt_date"`
	Total     string  `db:"total"`
}

func ListReceipts(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	log.Print(vars)
	userId := vars["userId"]
	limit := vars["limit"]
	offset := vars["offset"]

	if userId == "" {
		w.WriteHeader(http.StatusBadRequest)
		log.Printf("invalid receipt id")
		return
	}

	query := `SELECT * FROM public.receipt_data_v1 rd JOIN project_v1 p ON rd.project_id = p.project_id WHERE p.user_id = $1 ORDER BY rd.receipt_date DESC`
	if limit != "" {
		query += " LIMIT " + limit
	}
	if offset != "" {
		query += " OFFSET " + offset
	}

	db, ok := r.Context().Value("db").(*sqlx.DB)
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("db not found in context")
		return
	}

	var receipts []receiptResponse
	err := db.Select(&receipts, query, userId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("failed to get receipt snapshot: %s", err.Error())
		return
	}

	log.Printf("receipts: %#v", receipts)
	res, err := json.Marshal(receipts)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("json marshal error: %s", err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
	log.Printf("response body: %s", res)
}
