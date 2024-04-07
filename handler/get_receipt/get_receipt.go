package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type getReceiptRequest struct {
	ReceiptId uint64 `json:"receipt_id"`
}

type receiptResponse struct {
	ReceiptId uint64  `db:"receipt_id"`
	ScanId    *uint64 `db:"scan_id"`
	Subtotal  string  `db:"subtotal"`
	Tax       string  `db:"tax"`
	Gratuity  string  `db:"gratuity"`
	Date      string  `db:"receipt_date"`
	Total     string  `db:"total"`
}

func GetReceipt(w http.ResponseWriter, r *http.Request) {
	var reqBody getReceiptRequest

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	err := dec.Decode(&reqBody)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Printf("read body error: %s", err.Error())
		return
	}

	log.Printf("request body: %#v", reqBody)

	if reqBody.ReceiptId == 0 {
		log.Printf("receipt_id field is empty")
		http.Error(w, "empty fields", http.StatusUnprocessableEntity)
		return
	}

	db, ok := r.Context().Value("db").(*sqlx.DB)
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("db not found in context")
		return
	}

	var receipt receiptResponse
	err = db.Get(&receipt, `SELECT * FROM public.receipt_data_v1 WHERE receipt_id = $1 LIMIT 1`, reqBody.ReceiptId)
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
