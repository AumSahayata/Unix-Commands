package internal

import (
	_ "fmt"
	"path/filepath"
	"slices"
	"testing"
	"gotree/test"
)

func TestScanDir(t *testing.T) {
	testSet := test.CreateTestDirSet(t)
	res := ScanDir(testSet[0], -1, false, false, false, false)
	
	lines := test.NormalizeResults(res)

	if len(lines) != len(testSet)+2 {
		t.Fatal("test failed: Number of files and dirs does not match")
	}

	for i:=1;i<len(testSet);i++ {
		if !slices.Contains(lines, filepath.Base(testSet[i])) {
			t.Fatal("test failed: file/dir name not found in the output")
		}
	}
}

func TestOnlyDir(t *testing.T) {
	testSet := test.CreateTestDirSet(t)
	res := ScanDir(testSet[0], -1, false, false, false, true)

	set := testSet[:3]
	
	lines := test.NormalizeResults(res)

	if len(lines) != len(set)+2 {
		t.Fatal("test failed: Number of files and dirs does not match")
	}
	
	for i:=1;i<len(set);i++ {
		if !slices.Contains(lines, filepath.Base(set[i])) {
			t.Fatal("test failed: file/dir name not found in the output")
		}
	}
}

func TestLevels(t *testing.T) {
	testSet := test.CreateTestDirSet(t)
	res := ScanDir(testSet[0], 1, false, false, false, false)
	
	set := testSet[:3]
	set = append(set, testSet[4:]...)

	lines := test.NormalizeResults(res)

	if len(lines) != len(set)+2 {
		t.Fatal("test failed: Number of files and dirs does not match")
	}
	
	for i:=1;i<len(set);i++ {
		if !slices.Contains(lines, filepath.Base(set[i])) {
			t.Fatal("test failed: file/dir name not found in the output")
		}
	}
}

func TestSort(t *testing.T) {
	testSet := test.CreateTestDirSet(t)
	res := ScanDir(testSet[0], -1, true, false, false, false)
	
	set := [5]string{testSet[2], testSet[3], testSet[1], testSet[5], testSet[4]}

	lines := test.NormalizeResults(res)

	if len(lines) != len(set)+3 {
		t.Fatal("test failed: Number of files and dirs does not match")
	}

	for i:=1;i<len(set);i++ {
		if !slices.Contains(lines, filepath.Base(set[i])) {
			t.Fatal("test failed: file/dir name not found in the output")
		}
	}
}