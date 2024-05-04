package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type createScanInferenceRequest struct {
	ProjectId       *uint64                `json:"project_id"`
	ImageUrl        string                 `json:"image_url"`
	InferenceResult map[string]interface{} `json:"inference_result"`
}

func CreateScanInference(w http.ResponseWriter, r *http.Request) {
	var reqBody createScanInferenceRequest

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	err := dec.Decode(&reqBody)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Printf("read body error: %s", err.Error())
		return
	}

	log.Printf("request body: %#v", reqBody)

	if reqBody.ProjectId == nil || reqBody.InferenceResult == nil || reqBody.ImageUrl == "" {
		log.Printf("store scan request body contains empty fields")
		http.Error(w, "empty fields", http.StatusUnprocessableEntity)
		return
	}

	db, ok := r.Context().Value("db").(*sqlx.DB)
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("store scan db not found in context")
		return
	}

	tx, err := db.Beginx()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("store scan tx begin error: %s", err.Error())
		return
	}

	defer tx.Rollback()

	insertScanRes := tx.QueryRowxContext(r.Context(), `INSERT INTO public.scan_v1 (project_id, image_url, create_time) VALUES ($1, $2, NOW()) RETURNING scan_id`, reqBody.ProjectId, reqBody.ImageUrl)
	var scanId uint64
	err = insertScanRes.Scan(&scanId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("failed to insert into scan table: %s", err.Error())
		return
	}

	var inferenceId uint64
	inferenceResult, err := json.Marshal(reqBody.InferenceResult)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("store inference result json marshal error: %s", err.Error())
		return
	}

	row := tx.QueryRowxContext(r.Context(), `INSERT INTO public.scan_inference_v1 (scan_id, inference_result) VALUES ($1, $2) RETURNING scan_inference_id`, scanId, inferenceResult)
	err = row.Scan(&inferenceId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("insert scan_inference table error: %s", err.Error())
		return
	}

	err = tx.Commit()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("store scan tx commit error: %s", err.Error())
		return
	}

	res, err := json.Marshal(map[string]interface{}{"scan_id": scanId, "scan_inference_id": inferenceId})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("json marshal error: %s", err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
	log.Printf("response body: %s", res)
}
