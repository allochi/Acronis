package services

import (
	"io"

	"github.com/allochi/Acronis/models"
)

type ArchiveService struct {
}

func NewArchiveService() *ArchiveService {
	return &ArchiveService{}
}

func (s *ArchiveService) Write(w io.Writer, a models.Archive) error {
	return nil
}
