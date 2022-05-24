package api

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/allochi/Acronis/models"
	"github.com/allochi/Acronis/services"
	"github.com/google/uuid"
)

type DownloadAPI struct {
	downloads map[string]context.CancelFunc
}

func NewDownloadAPI(mux *http.ServeMux) {
	var api = &DownloadAPI{
		downloads: make(map[string]context.CancelFunc),
	}

	mux.Handle("/download", api)
	mux.Handle("/download/list", api)
	mux.Handle("/download/cancel", api)
}

func (api *DownloadAPI) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	method := r.Method
	log.Println(path)

	switch path {
	case "/download":
		switch method {
		case http.MethodPost:
			api.getArchive(w, r)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}

	case "/download/list":
		switch method {
		case http.MethodGet:
			api.list(w, r)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}

	case "/download/cancel":
		switch method {
		case http.MethodPost:
			api.cancel(w, r)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}

	}

}

func (api *DownloadAPI) list(w http.ResponseWriter, r *http.Request) {
	var keys []string
	for k := range api.downloads {
		keys = append(keys, k)
	}
	fmt.Fprintf(w, "`%v`\n", keys)
}

func (api *DownloadAPI) cancel(w http.ResponseWriter, r *http.Request) {
	var body, _ = ioutil.ReadAll(r.Body)
	var key = string(body)
	fmt.Fprintf(w, "canceling %s", key)
	api.downloads[key]()
}

func (api *DownloadAPI) getArchive(w http.ResponseWriter, r *http.Request) {
	var id = uuid.New().String()
	var ctx, cancel = context.WithCancel(r.Context())
	api.downloads[id] = cancel

	var files []models.File

	// Parse request
	err := json.NewDecoder(r.Body).Decode(&files)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// create an archive
	archive := models.NewArchive()
	archive.Add(files...)

	// create an archive service
	service := services.NewArchiveService(ctx, archive)
	_, err = service.Write(w)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}
