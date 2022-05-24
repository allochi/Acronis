package services

import (
	"archive/zip"
	"context"
	"fmt"
	"io"
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

// Write create a temporary archive zip file and add files to it
func (s *ArchiveService) Write(w io.Writer) (int, error) {
	// process files
	zipper := zip.NewWriter(w)
	for _, file := range s.archive.Files() {
		// handel request cancellation
		err := s.ctx.Err()
		if err != nil {
			log.Printf("process canceled: %s", err)
			return 0, err
		}

		// log.Printf("archiving: %s", file)

		err = s.processFile(zipper, file)
		if err != nil {
			log.Printf("failed to archive %s - %s", file, err)
			continue
		}

		// log.Printf("file %s archived", file)
	}
	zipper.Close()
	return 0, nil
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
