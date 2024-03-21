package create_user

import (
	"encoding/json"
	"net/http"

	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type userSnapshot struct {
	UserId      int64
	GivenName   string
	MiddleName  string
	FamilyName  string
	DateOfBirth string
	Email       string
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	db, err := sqlx.Connect("postgres", os.Getenv("POSTGRES_DSN"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	err = r.ParseForm()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	tx := db.MustBegin()

	insertRes := tx.MustExec(`INSERT INTO public.user_v1 (organization_id) VALUES $1 RETURNING user_id`, r.FormValue("organization_id"))
	userId, err := insertRes.LastInsertId()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	snapshot := userSnapshot{
		UserId:      userId,
		GivenName:   r.FormValue("given_name"),
		MiddleName:  r.FormValue("middle_name"),
		FamilyName:  r.FormValue("family_name"),
		DateOfBirth: r.FormValue("date_of_birth"),
		Email:       r.FormValue("email"),
	}
	tx.NamedExec(`INSERT INTO public.user_snapshot_v1 (user_id, given_name, middle_name, family_name, date_of_birth, email, create_time) VALUES ($1, $2, $3, $4, $5, $6, NOW())`, &snapshot)

	err = tx.Commit()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)

	res, err := json.Marshal(snapshot)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Write(res)
}
