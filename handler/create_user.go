package handler

import (
	"encoding/json"
	"io"
	"net/http"

	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Info().Ctx(ctx).Msgf("read body error: %s", err.Error())
		return
	}

	log.Info().Ctx(ctx).Msgf("request body: %s", body)

	w.Header().Set("Content-Type", "application/json")

	err = r.ParseForm()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Info().Ctx(ctx).Msgf("parsing form error: %s", err.Error())
		return
	}

	db, err := sqlx.Connect("postgres", os.Getenv("DSN"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Info().Ctx(ctx).Msgf("db connection error: %s", err.Error())
		w.Write([]byte("error: db not connected"))
		return
	}

	userId := r.FormValue("user_id")
	if userId == "" {
		w.WriteHeader(http.StatusBadRequest)
		log.Info().Ctx(ctx).Msgf("error: missing user_id field")
		w.Write([]byte("error: missing fields"))
		return
	}

	tx := db.MustBegin()
	defer tx.Rollback()

	var orgId int64
	insertOrgRes := tx.QueryRowxContext(ctx, `INSERT INTO public.organization_v1 DEFAULT VALUES RETURNING organization_id`)
	err = insertOrgRes.Scan(&orgId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Info().Ctx(ctx).Msgf("scan error: %s", err.Error())
		return
	}

	_, err = tx.QueryxContext(ctx, `INSERT INTO public.user_v1 (organization_id, user_id) VALUES ($1, $2) RETURNING user_id`, orgId, userId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Info().Ctx(ctx).Msgf("insert user error: %s", err.Error())
		return
	}

	err = tx.Commit()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Info().Ctx(ctx).Msgf("tx commit error: %s", err.Error())
		return
	}

	res, err := json.Marshal(map[string]string{"user_id": userId})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Info().Ctx(ctx).Msgf("json marshal error: %s", err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(res)
	log.Info().Ctx(ctx).Msgf("response body: %s", res)
}
