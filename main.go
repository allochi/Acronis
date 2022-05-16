package main

import (
	"log"
	"net/http"
	"os"

	"github.com/allochi/Acronis/api"
)

// TODO: File Handlers need to be managed within system boundaries
func main() {
	// // get request
	// request := []byte(`[
	//     "/Users/allochi/Projects/Acronis/sample-files/sample.yml",
	//     "/Users/allochi/Projects/Acronis/sample-files/sample.json",
	//     "/Users/allochi/Projects/Acronis/sample-files/sample.txt",
	//     "/Users/allochi/Projects/Acronis/sample-files/unavailable-file.txt"
	// ]`)

	// // parse request
	// var files []File
	// err := json.Unmarshal(request, &files)
	// if err != nil {
	// 	panic(err)
	// }

	// litter.Dump(files)

	// archive := NewArchive(files)
	// archive.Write()

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
