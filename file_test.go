package main

import (
	"os"
	"path/filepath"
	"testing"
)

var file File

func init() {
	wd, _ := os.Getwd()
	file = File(filepath.Join(wd, "file_test.go"))
}

func TestFileIsExist(t *testing.T) {
	if !file.IsExist() {
		t.Errorf("expected file %s to exist", file)
	}
}

func TestFileIsValid(t *testing.T) {
	var tests = []struct {
		file     File
		expected bool
	}{
		{"", false},
		{" ", false},
		{"/", false},
		{"//", false},
		{"/var/", false},
		{"/var", false},
		{"/dev/stdin", false},
		{file, true},
	}

	for _, test := range tests {
		if test.file.IsValid() != test.expected {
			t.Errorf("expected file %s .IsValid to be %t", test.file, test.expected)
		}
	}

}
