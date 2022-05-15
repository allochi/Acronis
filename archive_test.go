package main

import "testing"

func TestAddFile(t *testing.T) {
	files := []File{
		"/home/user/log.txt",
		"/home/user/messages.txt",
	}

	archive := NewArchive(files)
	archive.Add("/home/user/messages.txt")

	// check for duplication prevention
	exp := len(files)
	got := len(archive.files)
	if exp != got {
		t.Errorf("file duplicated, expected %d files got %d files", exp, got)
	}

	// check for uniqueness
	archive.Add("/home/user/config.txt")
	exp = len(files) + 1
	got = len(archive.files)
	if exp != got {
		t.Errorf("file duplicated, expected %d files got %d files", exp, got)
	}
}
