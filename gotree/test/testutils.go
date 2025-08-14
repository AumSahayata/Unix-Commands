package test

import (
	"os"
	"strings"
	"testing"
)

func CreateTestDirSet(t *testing.T) [6]string {
	t.Helper()

	var testSet [6]string

	tempDir, err := os.MkdirTemp("", "tree-test-dir")
	if err != nil {
		t.Fatal("Failed to create temp dir:", err)
	}

	testSet[0] = tempDir

	f, err := os.CreateTemp(tempDir, "file-1.txt")
	if err != nil {
		t.Fatal("Failed to create temp file in dir:", err)
	}

	testSet[4] = f.Name()

	f, err = os.CreateTemp(tempDir, "file-2.txt")
	if err != nil {
		t.Fatal("Failed to create temp file in dir:", err)
	}

	testSet[5] = f.Name()

	d, err := os.MkdirTemp(tempDir, "d1")
	if err != nil {
		t.Fatal("Failed to create temp dir in dir:", err)
	}

	testSet[1] = d

	f, err = os.CreateTemp(d, "in-dir.txt")
	if err != nil {
		t.Fatal("Failed to create temp file in dir:", err)
	}

	testSet[3] = f.Name()

	d, err = os.MkdirTemp(tempDir, "d2")
	if err != nil {
		t.Fatal("Failed to create temp dir in dir:", err)
	}

	testSet[2] = d

	t.Cleanup(func() {
		os.RemoveAll(tempDir)
	})

	return testSet
}

func NormalizeResults(in string) []string {
	in = strings.ReplaceAll(in, " ", "")
	in = strings.ReplaceAll(in, "└──", "")
	in = strings.ReplaceAll(in, "│", "")
	in = strings.ReplaceAll(in, "├──", "")
	lines := strings.Split(strings.TrimSpace(in), "\n")

	return lines
}