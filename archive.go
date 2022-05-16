package main

import (
	"archive/zip"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

// Archive a data structure that holds a set of unique files
// and compile them into one zip file
type Archive struct {
	files map[File]struct{}
}

// NewArchive create new archive of files
func NewArchive(files []File) *Archive {
	archive := Archive{
		files: make(map[File]struct{}),
	}

	// Add valid files to the files sets
	for _, file := range files {
		err := archive.Add(file)
		if err != nil {
			// TODO: may report list of errors back the request
			log.Println(err)
		}
	}

	return &archive
}

// Add add unique valid files to the archive
func (a *Archive) Add(file File) error {
	if !file.IsValid() {
		return fmt.Errorf("failed to archive %s - invalid file", file)
	}
	a.files[file] = struct{}{}
	return nil
}

// Write create a temproray archive zip file and add files to it
func (a *Archive) Write() (int, error) {
	// WARN: How big this file could reach, and how to handle system capacity
	archive, err := ioutil.TempFile("/tmp", "archive.*.zip")
	if err != nil {
		return 0, err
	}
	defer os.Remove(archive.Name())

	zipper := zip.NewWriter(archive)
	defer zipper.Close()

	// TODO: goroutines zip files
	for file := range a.files {
		log.Printf("archiving %s...", file)

		err = a.zipFile(zipper, file)
		if err != nil {
			log.Printf("failed to archive %s - ", err)
			continue
		}

		log.Printf("file %s archived", file)
	}

	return 0, nil
}

// zipFile handles adding one file to the archive
func (a *Archive) zipFile(zipper *zip.Writer, file File) error {
	f, err := os.Open(file.String())
	if err != nil {
		return fmt.Errorf("can't open file")
	}
	defer f.Close()

	w, err := zipper.Create(file.String())
	if err != nil {
		return fmt.Errorf("can't zip file")
	}

	_, err = io.Copy(w, f)
	if err != nil {
		return fmt.Errorf("can't write file to archive")
	}

	return nil
}
