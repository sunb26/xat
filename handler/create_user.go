package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	db, err := sqlx.Connect("postgres", os.Getenv("POSTGRES_DSN"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("db connection error: %s", err.Error())))
		return
	}

	err = r.ParseForm()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("parse form error: %s", err.Error())))
		return
	}

	userId := r.FormValue("user_id")
	if userId == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("error: missing user_id"))
		return
	}

	tx := db.MustBegin()

	var orgId int64
	insertOrgRes := tx.QueryRowx(`INSERT INTO public.organization_v1 DEFAULT VALUES RETURNING organization_id`)
	err = insertOrgRes.Scan(&orgId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("scan error: %s", err.Error())))
		return
	}

	tx.MustExec(`INSERT INTO public.user_v1 (organization_id, user_id) VALUES ($1, $2) RETURNING user_id`, orgId, userId)

	err = tx.Commit()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("tx commit error: %s", err.Error())))
		return
	}

	res, err := json.Marshal(map[string]string{"user_id": userId})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("json marshal error: %s", err.Error())))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
