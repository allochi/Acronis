package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestFileExists(t *testing.T) {
	wd, _ := os.Getwd()
	var file = File(filepath.Join(wd, "file_test.go"))

	if !file.Exists() {
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
		{"/var/file.yml", true},
	}

	for _, test := range tests {
		if test.file.IsValid() != test.expected {
			t.Errorf("expected file %s .IsValid to be %t", test.file, test.expected)
		}
	}

}
