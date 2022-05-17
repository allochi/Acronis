package models

import (
	"os"
	"path/filepath"
	"testing"
)

func TestAddFile(t *testing.T) {
	wd, _ := os.Getwd()
	file1 := File(filepath.Join(wd, "archive.go"))
	file2 := File(filepath.Join(wd, "archive_test.go"))

	archive := NewArchive()
	archive.Add(file1, file2)

	// check for duplication prevention
	file3 := File(filepath.Join(wd, "archive.go"))
	archive.Add(file3)
	exp := 2
	got := len(archive.Files())
	if exp != got {
		t.Errorf("file duplicated, expected %d files got %d files", exp, got)
	}
}

func TestIsValid(t *testing.T) {
	archive := NewArchive()

	wd, _ := os.Getwd()
	file := File(filepath.Join(wd, "archive_test.go"))
	unavailable := File(filepath.Join(wd, "unavailable_test.go"))

	var tests = []struct {
		file     File
		expected bool
	}{
		{"", false},
		{" ", false},
		{"home/file.txt", false},
		{"/", false},
		{"//", false},
		{"/var/", false},
		{"/var", false},
		{"/dev/stdin", false},
		{unavailable, false},
		{file, true},
	}

	for _, test := range tests {
		if archive.IsValid(test.file) != test.expected {
			t.Errorf("expected file %s .IsValid to be %t", test.file, test.expected)
		}
	}

}
