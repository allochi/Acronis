package models

import (
	"archive/zip"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"time"
)

// Archive a data structure that holds a set of unique files
// and compile them into one zip file
type Archive struct {
	files map[File]struct{}
	// failed map[File]struct{}
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
			// archive.failed[file] = struct{}{}
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
// WARN: How big this file could reach, and how to handle system capacity
func (a *Archive) Write(w io.Writer) (int, error) {
	// create archive file
	archive, err := ioutil.TempFile("/tmp", "archive.*.zip")
	if err != nil {
		return 0, err
	}
	defer os.Remove(archive.Name())

	// create a zip writer
	zipper := zip.NewWriter(archive)

	start := time.Now()
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
	log.Printf("duration %v", time.Since(start))
	// TODO report failed files
	// zipper.SetComment("What is this?")
	zipper.Close()

	// write archive file to external writer
	// TODO do we need this???
	file, err := os.Open(archive.Name())
	if err != nil {
		log.Panicln(err)
	}
	n, err := io.Copy(w, file)
	log.Println(n)
	return 0, err
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
