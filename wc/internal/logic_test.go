package internal

import (
	"testing"
	"wc/test"
)

func TestCountLines(t *testing.T) {
	file := test.CreateTestFile(t)
	
	f, err := OpenFile(file.Name())
	if err != nil {
		t.Fatal("Failed to open file.")
	}
	defer f.Close()

	lines, err := CountLines(f)
	if err != nil {
		t.Fatal("Failed to count lines in file:", err)
	}
	
	if lines != 3 {
		t.Fatal("Miscounted the number of lines in the file.")
	}
}

func TestCountWords(t *testing.T) {
	file := test.CreateTestFile(t)

	f, err := OpenFile(file.Name())
	if err != nil {
		t.Fatal("Failed to open file.")
	}
	defer f.Close()
	
	words, err := CountWords(f)
	if err != nil {
		t.Fatal("Failed to count words in file:", err)
	}

	if words != 5 {
		t.Fatal("Miscounted the number of words in the file.")
	}
}

func TestCountCharacters(t *testing.T) {
	file := test.CreateTestFile(t)

	f, err := OpenFile(file.Name())
	if err != nil {
		t.Fatal("Failed to open file.")
	}
	defer f.Close()

	chars, err := CountCharacters(f)
		if err != nil {
		t.Fatal("Failed to count characters in file:", err)
	}

	if chars != 24 {
		t.Fatal("Miscounted the number of characters in the file.")
	}
}