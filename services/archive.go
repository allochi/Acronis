package services

import (
	"archive/zip"
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"

	"github.com/allochi/Acronis/models"
)

type ArchiveService struct {
	ctx     context.Context
	archive *models.Archive
}

func NewArchiveService(ctx context.Context, archive *models.Archive) *ArchiveService {
	return &ArchiveService{
		ctx:     ctx,
		archive: archive,
	}
}

// Write create a temproray archive zip file and add files to it
func (s *ArchiveService) Write(w io.Writer) (int, error) {
	// create archive file
	// WARN: need to handel file size and system capacity
	archive, err := ioutil.TempFile("/tmp", "archive.*.zip")
	if err != nil {
		return 0, err
	}
	defer os.Remove(archive.Name())

	// process files
	zipper := zip.NewWriter(archive)
	for _, file := range s.archive.Files() {
		// handel request cancelation
		err := s.ctx.Err()
		if err != nil {
			log.Printf("process canceld: %s", err)
			return 0, err
		}

		log.Printf("archiving: %s", file)

		err = s.processFile(zipper, file)
		if err != nil {
			log.Printf("failed to archive %s - %s", file, err)
			continue
		}

		log.Printf("file %s archived", file)
	}
	zipper.Close()

	// write archive file to external writer
	archive.Seek(0, 0)
	_, err = io.Copy(w, archive)
	return 0, err
}

// processFile handles adding one file to the archive
func (s *ArchiveService) processFile(zipper *zip.Writer, file models.File) error {
	path := string(file)

	f, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("can't open file %s", file)
	}
	defer f.Close()

	w, err := zipper.Create(path)
	if err != nil {
		return fmt.Errorf("can't zip file %s", file)
	}

	_, err = io.Copy(w, f)
	if err != nil {
		return fmt.Errorf("can't write to archive file %s", file)
	}

	return nil
}
