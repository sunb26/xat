package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type createReceiptRequest struct {
	ProjectId string `json:"project_id"`
	Tax       string `json:"tax"`
	Gratuity  string `json:"gratuity"`
	Date      string `json:"date"`
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

	if reqBody.ProjectId == "" || reqBody.Tax == "" || reqBody.Gratuity == "" || reqBody.Date == "" {
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

	insertReceiptRes := tx.QueryRowxContext(r.Context(), `INSERT INTO public.receipt_v1 VALUES (project_id) RETURNING receipt_id`)
	var receiptId uint64
	err = insertReceiptRes.Scan(&receiptId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("failed to insert into receipt table: %s", err.Error())
		return
	}

	_, err = tx.QueryxContext(r.Context(), `INSERT INTO public.receipt_snapshot_v1 (receipt_id, receipt_date, create_time) VALUES ($1, $2, NOW()) RETURNING receipt_id`, receiptId, reqBody.Date)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("insert receipt error: %s", err.Error())
		return
	}

	ids, err := tx.QueryxContext(r.Context(), `INSERT INTO public.expense_v1 (receipt_id) VALUES ($1), ($1) RETURNING expense_id`, receiptId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("insert receipt error: %s", err.Error())
		return
	}

	var expenseId uint64

	err = ids.Scan(&expenseId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("scan tax expense id error: %s", err.Error())
		return
	}
	_, err = tx.QueryxContext(r.Context(), `INSERT INTO public.expense_snapshot_v1 (expense_id, tag, title, amount, deductible, create_time) VALUES ($1, "tax", "GST/HST", $2, 0, NOW())`, expenseId, reqBody.Tax)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("insert tax expense error: %s", err.Error())
		return
	}

	ids.Next()
	err = ids.Scan(&expenseId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("scan gratuity expense id error: %s", err.Error())
		return
	}

	_, err = tx.QueryxContext(r.Context(), `INSERT INTO public.expense_snapshot_v1 (expense_id, tag, title, amount, deductible, create_time) VALUES ($1, "gratuity", "Gratuity", $2, 0, NOW())`, expenseId, reqBody.Gratuity)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("insert gratuity expense error: %s", err.Error())
		return
	}

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
