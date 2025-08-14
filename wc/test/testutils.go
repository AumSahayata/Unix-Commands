package test

import (
	"os"
	"testing"
)

func CreateTestFile(t *testing.T) *os.File {
	t.Helper()

	tempFile, err := os.CreateTemp("", "-wc-test-*.txt")
	if err != nil {
		t.Fatal("Failed to create temp test file:", err)
	}

	_, err = tempFile.WriteString("Hello World\nTesting\n3rd line")
	if err != nil {
		t.Fatal("Failed to write in temp test file:", err)
	}

	t.Cleanup(func() {
		tempFile.Close()
		os.Remove(tempFile.Name())
	})

	return tempFile
}