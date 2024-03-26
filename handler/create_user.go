package handler

import (
	"encoding/json"
	"net/http"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"
)

type createUserRequest struct {
	UserId string `json:"user_id"`
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var reqBody createUserRequest

	// TODO: Put error handling in helper function
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	err := dec.Decode(&reqBody)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Info().Ctx(ctx).Msgf("read body error: %s", err.Error())
		return
	}

	log.Info().Ctx(ctx).Msgf("request body: %#v", reqBody)

	w.Header().Set("Content-Type", "application/json")

	if reqBody.UserId == "" {
		log.Info().Ctx(ctx).Msg("user_id field is empty")
		http.Error(w, "empty fields", http.StatusUnprocessableEntity)
		return
	}

	db, ok := r.Context().Value("db").(*sqlx.DB)
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		log.Info().Ctx(ctx).Msg("db not found in context")
		return
	}

	tx, err := db.Beginx()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Info().Ctx(ctx).Msgf("tx begin error: %s", err.Error())
		return
	}

	defer tx.Rollback()

	var orgId uint64
	insertOrgRes := tx.QueryRowxContext(ctx, `INSERT INTO public.organization_v1 DEFAULT VALUES RETURNING organization_id`)
	err = insertOrgRes.Scan(&orgId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Info().Ctx(ctx).Msgf("failed to insert into organization table: %s", err.Error())
		return
	}

	_, err = tx.QueryxContext(ctx, `INSERT INTO public.user_v1 (organization_id, user_id) VALUES ($1, $2) RETURNING user_id`, orgId, reqBody.UserId)
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

	res, err := json.Marshal(map[string]string{"user_id": reqBody.UserId})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Info().Ctx(ctx).Msgf("json marshal error: %s", err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(res)
	log.Info().Ctx(ctx).Msgf("response body: %s", res)
}
