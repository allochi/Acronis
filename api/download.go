package api

import (
	"encoding/json"
	"net/http"

	"github.com/allochi/Acronis/models"
	"github.com/allochi/Acronis/services"
)

type DownloadAPI struct {
	service *services.ArchiveService
}

func NewDownloadAPI() *DownloadAPI {
	return &DownloadAPI{
		service: services.NewArchiveService(),
	}
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
	archive := models.NewArchive(files)
	archive.Write(w)

	// w.WriteHeader(http.StatusOK)
	// pass the archive and a writer to the service
	// finalize the response

	// products := api.service.GetProducts()
	// err := api.service.EncodeSliceJSON(w, products)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }
}
