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
	ReceiptId uint64  `db:"receipt_id"`
	ScanId    *uint64 `db:"scan_id"`
	Subtotal  string  `db:"subtotal"`
	Tax       string  `db:"tax"`
	Gratuity  string  `db:"gratuity"`
	Date      string  `db:"receipt_date"`
	Total     string  `db:"total"`
}

func GetReceipt(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	log.Print(vars)
	receiptId := vars["receiptId"]

	if receiptId == "" {
		w.WriteHeader(http.StatusBadRequest)
		log.Printf("invalid receipt id")
		return
	}

	db, ok := r.Context().Value("db").(*sqlx.DB)
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("db not found in context")
		return
	}

	var receipt receiptResponse
	err := db.Get(&receipt, `SELECT * FROM public.receipt_data_v1 WHERE receipt_id = $1 LIMIT 1`, receiptId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("failed to get receipt snapshot: %s", err.Error())
		return
	}

	res, err := json.Marshal(receipt)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("json marshal error: %s", err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
	log.Printf("response body: %s", res)
}
