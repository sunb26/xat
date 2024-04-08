package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type createReceiptRequest struct {
	ProjectId *uint64 `json:"project_id"`
	Tax       string  `json:"tax"`
	Gratuity  string  `json:"gratuity"`
	Date      string  `json:"date"`
	ScanId    *uint64 `json:"scan_id"`
}

func CreateReceipt(w http.ResponseWriter, r *http.Request) {
	var reqBody createReceiptRequest

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	err := dec.Decode(&reqBody)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Printf("read body error: %s", err.Error())
		return
	}

	log.Printf("request body: %#v", reqBody)

	if reqBody.ProjectId == nil || reqBody.Tax == "" || reqBody.Gratuity == "" || reqBody.Date == "" {
		log.Printf("request body contains empty fields")
		http.Error(w, "empty fields", http.StatusUnprocessableEntity)
		return
	}

	db, ok := r.Context().Value("db").(*sqlx.DB)
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("db not found in context")
		return
	}

	tx, err := db.Beginx()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("tx begin error: %s", err.Error())
		return
	}

	defer tx.Rollback()

	insertReceiptRes := tx.QueryRowxContext(r.Context(), `INSERT INTO public.receipt_v1 (project_id) VALUES ($1) RETURNING receipt_id`, reqBody.ProjectId)
	var receiptId uint64
	err = insertReceiptRes.Scan(&receiptId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("failed to insert into receipt table: %s", err.Error())
		return
	}

	rows, err := tx.QueryxContext(r.Context(), `INSERT INTO public.receipt_snapshot_v1 (receipt_id, scan_id, receipt_date, create_time) VALUES ($1, $2, $3, NOW())`, receiptId, reqBody.ScanId, reqBody.Date)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("insert receipt_snapshot table error: %s", err.Error())
		return
	}
	rows.Close()

	ids, err := tx.QueryxContext(r.Context(), `INSERT INTO public.expense_v1 (receipt_id) VALUES ($1), ($1) RETURNING expense_id`, receiptId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("insert expense table error: %s", err.Error())
		return
	}
	var expenseIds []uint64

	for ids.Next() {
		var expenseId uint64
		err = ids.Scan(&expenseId)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Printf("scan tax and gratuity expense ids error: %s", err.Error())
			return
		}
		expenseIds = append(expenseIds, expenseId)
	}

	rows, err = tx.QueryxContext(r.Context(), `INSERT INTO public.expense_snapshot_v1 (expense_id, scan_id, tag, title, amount, deductible, create_time) VALUES ($1, $2, 'tax', 'GST/HST', $3, 0, NOW())`, expenseIds[0], reqBody.ScanId, reqBody.Tax)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("insert tax expense error: %s", err.Error())
		return
	}
	rows.Close()

	rows, err = tx.QueryxContext(r.Context(), `INSERT INTO public.expense_snapshot_v1 (expense_id, scan_id, tag, title, amount, deductible, create_time) VALUES ($1, $2, 'gratuity', 'Gratuity', $3, 0, NOW())`, expenseIds[1], reqBody.ScanId, reqBody.Gratuity)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("insert gratuity expense error: %s", err.Error())
		return
	}
	rows.Close()

	err = tx.Commit()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("tx commit error: %s", err.Error())
		return
	}

	res, err := json.Marshal(map[string]uint64{"receipt_id": receiptId})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("json marshal error: %s", err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
	log.Printf("response body: %s", res)
}
