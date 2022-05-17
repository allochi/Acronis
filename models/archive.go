package models

import (
	"os"
	"path/filepath"
	"strings"
)

// File represent a file path
type File string

// Archive a data structure that holds a set of unique files
type Archive struct {
	files       map[File]struct{}
	isPermitted func(File) bool
}

// NewArchive create new archive of files
func NewArchive() *Archive {
	return &Archive{
		files:       make(map[File]struct{}),
		isPermitted: func(f File) bool { return true },
	}
}

// Add add unique valid files to the archive
func (a *Archive) Add(files ...File) {
	for _, file := range files {
		if a.IsValid(file) {
			a.files[file] = struct{}{}
		}
	}
}

// Files return all files in the archive
func (a *Archive) Files() []File {
	var files = make([]File, 0, len(a.files))
	for file := range a.files {
		files = append(files, file)
	}
	return files
}

// IsValid
func (a *Archive) IsValid(file File) bool {
	var path = string(file)

	// not empty
	if strings.TrimSpace(path) == "" {
		return false
	}

	// only absolute path
	if !filepath.IsAbs(path) {
		return false
	}

	// only files
	stat, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	if stat.IsDir() {
		return false
	}
	if !stat.Mode().IsRegular() {
		return false
	}

	// only permitted
	if !a.isPermitted(file) {
		return false
	}

	return true
}

// SetPermissionRules set a function with validation rules
// to checks if a file can be retrived by the user
func (a *Archive) SetPermissionRules(isPermitted func(File) bool) {
	a.isPermitted = isPermitted
}
