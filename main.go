package main

import (
	"log"
	"net/http"
	"os"

	"github.com/allochi/Acronis/api"
)

func main() {
	pid := os.Getpid()

	mux := http.NewServeMux()
	api.NewDownloadAPI(mux)

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
