package api

import (
	"encoding/json"
	"net/http"

	"github.com/allochi/Acronis/models"
	"github.com/allochi/Acronis/services"
)

type DownloadAPI struct{}

func NewDownloadAPI() *DownloadAPI {
	return &DownloadAPI{}
}

func (api *DownloadAPI) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {

	case http.MethodPost:
		api.getArchive(w, r)

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (api *DownloadAPI) getArchive(w http.ResponseWriter, r *http.Request) {
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
	service := services.NewArchiveService(r.Context(), archive)
	_, err = service.Write(w)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}
