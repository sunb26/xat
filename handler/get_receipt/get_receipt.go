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

type snapshot struct {
	ReceiptId   uint64  `db:"receipt_id"`
	SnapshotId  uint64  `db:"snapshot_id"`
	ScanId      *uint64 `db:"scan_id"`
	Tax         string  `db:"tax"`
	Gratuity    string  `db:"gratuity"`
	ReceiptDate string  `db:"receipt_date"`
	CreateTime  string  `db:"create_time"`
}

type receiptResponse struct {
	ReceiptId uint64
	ScanId    uint64
	Subtotal  string
	Tax       string
	Gratuity  string
	Date      string
	Total     string
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

	var receiptSnapshot snapshot
	var subtotal string
	var total string

	err = db.Get(&receiptSnapshot, `SELECT * FROM public.receipt_snapshot_v1 WHERE receipt_id = $1 LIMIT 1`, reqBody.ReceiptId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("failed to get receipt snapshot: %s", err.Error())
		return
	}

	err = db.Get(&subtotal, `SELECT SUM(es.amount) AS subtotal FROM expense_v1 AS e JOIN expense_snapshot_v1 AS es ON e.expense_id = es.expense_id WHERE e.receipt_id = $1 GROUP BY e.receipt_id`, reqBody.ReceiptId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("failed to get subtotal: %s", err.Error())
		return
	}

	err = db.Get(&total, `SELECT COALESCE(SUM(e.amount), '$0'::money) + COALESCE(rs.gratuity, '$0'::money) + COALESCE(rs.tax, '$0'::money) AS total_amount FROM receipt_v1 r LEFT JOIN expense_v1 ex ON r.receipt_id = ex.receipt_id LEFT JOIN expense_snapshot_v1 e ON ex.expense_id = e.expense_id LEFT JOIN receipt_snapshot_v1 rs ON r.receipt_id = rs.receipt_id WHERE r.receipt_id = $1 GROUP BY rs.gratuity, rs.tax`, reqBody.ReceiptId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("failed to get subtotal: %s", err.Error())
		return
	}

	// handle NULL scan_id values
	var scanId uint64
	if receiptSnapshot.ScanId != nil {
		scanId = *receiptSnapshot.ScanId
	}

	receipt := receiptResponse{
		ReceiptId: reqBody.ReceiptId,
		ScanId:    scanId,
		Subtotal:  subtotal,
		Tax:       receiptSnapshot.Tax,
		Gratuity:  receiptSnapshot.Gratuity,
		Date:      receiptSnapshot.ReceiptDate,
		Total:     total,
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
