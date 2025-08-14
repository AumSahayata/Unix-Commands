package test

import (
	"os"
	"testing"
)

func CreateTestFile(t *testing.T) *os.File {
	t.Helper()

	file, err := os.CreateTemp("", "-grep-test-file-*")
    if err != nil {
        t.Fatal("Failed to create temp file:", err)
    }

    _, err = file.WriteString("098 starting\nHello World\nTesting hello\n123 closing")
    if err != nil {
        t.Fatal("Failed to write in temp file:", err)
    }

	t.Cleanup(func() {
		file.Close()
		os.Remove(file.Name())
	})

	return file
}