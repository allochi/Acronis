package main

import (
	"archive/zip"
	"io"
	"io/ioutil"
	"os"
)

// Archive a data structure that holds a set of unique files
// and compile them into one zip file
type Archive struct {
	// file  *os.File
	files map[File]struct{}
}

// NewArchive create new archive of files
func NewArchive(files []File) *Archive {
	archive := Archive{
		files: make(map[File]struct{}),
	}

	for _, file := range files {
		archive.files[file] = struct{}{}
	}

	return &archive
}

// Add add a file to the archive
func (a *Archive) Add(file File) {
	a.files[file] = struct{}{}
}

func (a *Archive) Write() (int, error) {
	archive, err := ioutil.TempFile("/tmp", "archive.*.zip")
	if err != nil {
		return 0, err
	}
	defer os.Remove(archive.Name())

	zipper := zip.NewWriter(archive)
	defer zipper.Close()

	for file := range a.files {
		f, err := os.Open(file.String())
		if err != nil {
			return 0, nil
		}

		w, err := zipper.Create(file.String())
		if err != nil {
			return 0, nil
		}
		if _, err := io.Copy(w, f); err != nil {
			return 0, nil
		}

		f.Close()
	}

	return 0, nil
}
