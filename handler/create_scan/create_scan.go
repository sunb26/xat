package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"

	"golang.org/x/oauth2/google"
	"golang.org/x/oauth2/jwt"
	"google.golang.org/api/drive/v3"
	"google.golang.org/api/option"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type createScanRequest struct {
	ProjectId *uint64 `json:"project_id"`
	ImageUrl  string  `json:"image_url"`
}

const fileIdPattern = `https:\/\/drive\.google\.com\/file\/d\/([a-zA-Z0-9_-]+)\/`

func CreateScan(w http.ResponseWriter, r *http.Request) {
	var reqBody createScanRequest

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	err := dec.Decode(&reqBody)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Printf("read body error: %s", err.Error())
		return
	}

	log.Printf("request body: %#v", reqBody)

	if reqBody.ProjectId == nil || reqBody.ImageUrl == "" {
		log.Printf("store scan request body contains empty fields")
		http.Error(w, "empty fields", http.StatusUnprocessableEntity)
		return
	}

	err = verifyFileId(r.Context(), reqBody.ImageUrl)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("verify file id error: %s", err.Error())
		return
	}

	db, ok := r.Context().Value("db").(*sqlx.DB)
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("store image url: db not found in context")
		return
	}

	tx, err := db.Beginx()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("store image url: tx begin error: %s", err.Error())
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

	err = tx.Commit()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("store scan tx commit error: %s", err.Error())
		return
	}

	res, err := json.Marshal(map[string]interface{}{"scan_id": scanId, "project_id": *reqBody.ProjectId, "image_url": reqBody.ImageUrl})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("json marshal error: %s", err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
	log.Printf("response body: %s", res)
}

func verifyFileId(ctx context.Context, url string) error {

	re := regexp.MustCompile(fileIdPattern)
	match := re.FindStringSubmatch(url)

	var fileId string
	if len(match) > 1 {
		fileId = match[1]
		log.Printf("found file id: %s", fileId)
	} else {
		return fmt.Errorf("could not parse file id from url")
	}

	conf := &jwt.Config{
		Email:      os.Getenv("GOOGLE_SERVICE_EMAIL"),
		PrivateKey: []byte(os.Getenv("GOOGLE_PRIVATE_KEY")),
		Scopes: []string{
			"https://www.googleapis.com/auth/drive",
		},
		TokenURL: google.JWTTokenURL,
	}

	client := conf.Client(ctx)

	service, err := drive.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		return fmt.Errorf("creating drive service: %s", err)
	}

	found, err := service.Permissions.List(fileId).Do()
	if err != nil {
		return fmt.Errorf("listing gdrive file permissions: %s", err)
	}

	for _, perm := range found.Permissions {
		log.Printf("found permission: %#v", perm)
		if perm.EmailAddress == os.Getenv("GOOGLE_SERVICE_EMAIL") && perm.Role == "owner" {
			return nil
		}
	}

	return fmt.Errorf("invalid gdrive fileid")
}
