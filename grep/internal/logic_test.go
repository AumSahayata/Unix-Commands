package internal

import (
	"grep/test"
	"testing"
)

func TestSearchFileForZeroMatch(t *testing.T) {
	file := test.CreateTestFile(t)
	searchString := "Duck"
	res := SearchFile(searchString, file, false, 0, 0)

	if len(res) != 0 {
		t.Error("The number of matches is not correct.")
	}
}

func TestSearchFileForOneMatch(t *testing.T) {
	file := test.CreateTestFile(t)
	searchString := "123"
	res := SearchFile(searchString, file, false, 0, 0)

	if len(res) != 1 {
		t.Error("The number of matches is not correct.")
	}
}

func TestSearchFileForTwoMatch(t *testing.T) {
	file := test.CreateTestFile(t)
	searchString := "Hello"
	res := SearchFile(searchString, file, true, 0, 0)

	if len(res) != 2 {
		t.Error("The number of matches is not correct.")
	}
}

func TestNlinesBefore(t *testing.T) {
	file := test.CreateTestFile(t)
	searchString := "Test"
	res := SearchFile(searchString, file, true, 0, 2)

	// There are 2 lines before the searchString in the test file
	if len(res) != 3 {
		t.Error("The number of returnd lines is not correct.")
	}
}

func TestNlinesAfter(t *testing.T) {
	file := test.CreateTestFile(t)
	searchString := "Test"
	res := SearchFile(searchString, file, true, 2, 0)

	// There is only 1 line after the searchString in the test file
	if len(res) != 2 {
		t.Error("The number of returnd lines is not correct.")
	}
}