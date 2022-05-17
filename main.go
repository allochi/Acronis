package main

import (
	"log"
	"net/http"
	"os"

	"github.com/allochi/Acronis/api"
)

func main() {
	pid := os.Getpid()

	downloadAPI := api.NewDownloadAPI()

	mux := http.NewServeMux()
	mux.Handle("/downloadzip", downloadAPI)

	srv := &http.Server{
		Addr:    ":9000",
		Handler: mux,
	}

	log.Printf("process %d listening on %s", pid, srv.Addr)
	err := srv.ListenAndServe()
	if err != nil {
		panic(err)
	}

}
