package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type createUserRequest struct {
	UserId string `json:"user_id"`
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var reqBody createUserRequest

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	err := dec.Decode(&reqBody)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Printf("read body error: %s", err.Error())
		return
	}

	log.Printf("request: %s %s %#v", r.Method, r.URL, reqBody)

	if reqBody.UserId == "" {
		log.Printf("user_id field is empty")
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

	var orgId uint64
	insertOrgRes := tx.QueryRowxContext(r.Context(), `INSERT INTO public.organization_v1 DEFAULT VALUES RETURNING organization_id`)
	err = insertOrgRes.Scan(&orgId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("failed to insert into organization table: %s", err.Error())
		return
	}

	_, err = tx.QueryxContext(r.Context(), `INSERT INTO public.user_v1 (organization_id, user_id) VALUES ($1, $2) RETURNING user_id`, orgId, reqBody.UserId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("insert user error: %s", err.Error())
		return
	}

	err = tx.Commit()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("tx commit error: %s", err.Error())
		return
	}

	res, err := json.Marshal(map[string]string{"user_id": reqBody.UserId})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("json marshal error: %s", err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
	log.Printf("response body: %s", res)
}
