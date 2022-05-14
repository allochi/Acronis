package main

import (
	"os"
	"path/filepath"
)

type File string

// Exists check if the file exists
func (f *File) Exists() bool {
	_, err := os.Stat(f.String())
	return !os.IsNotExist(err)
}

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

// IsValid checks if the file path is valid
func (f File) IsValid() bool {
	var valid = true
	var path = f.String()
	valid = valid && filepath.IsAbs(path)
	return valid
}
