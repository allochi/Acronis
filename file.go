package main

import (
	"os"
	"path/filepath"
)

type File string

// Write implement writer interface to stream the file in binary format
// TODO: parity check and encryption check?
// TODO: What happen when the file is larger than the memory
func (f *File) Write(p []byte) (n int, err error) {
	return 0, nil
}

// String return the file path
func (f File) String() string {
	return string(f)
}

// Binary returns the binary representation of a file content
func (f File) Binary() []byte {
	return []byte{}
}

// Exists check if the file exists
func (f *File) IsExist() bool {
	_, err := os.Stat(f.String())
	return !os.IsNotExist(err)
}

// IsValid checks if the file is valid
// TODO: what about file handlers in unix/linux?
func (f File) IsValid() bool {
	var path = f.String()

	if !filepath.IsAbs(path) {
		return false
	}

	if !f.IsExist() {
		return false
	}

	stat, err := os.Stat(path)
	if err != nil {
		return false
	}
	if stat.IsDir() {
		return false
	}
	if !stat.Mode().IsRegular() {
		return false
	}

	return true
}

// IsReadPermitted checks if the user has read permission on a file
// TODO: This requires authorization system beyond the scope of this exercise
func (f File) IsReadPermitted() bool {
	return true
}
