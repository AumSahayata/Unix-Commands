package internal

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func OpenFile(filepath string) (*os.File, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func SearchFile(searchString string, rd io.Reader, sens bool, a, b int) []string {
	if err := resetReader(rd); err != nil {
		return []string{}
	}

	nlines := make([]string, 0, b)

	var foundStrings []string
	scanner := bufio.NewScanner(rd)

	for scanner.Scan() {
		mainString := scanner.Text()

		if sens {
			mainString = strings.ToLower(scanner.Text())
			searchString = strings.ToLower(searchString)
		}

		if strings.Contains(mainString, searchString) {
			if b > 0 {
				foundStrings = append(foundStrings, nlines...)
			}
			foundStrings = append(foundStrings, scanner.Text())
			if a > 0 {
				foundStrings = append(foundStrings, getLinesAfter(scanner, a)...)
			}
		}

		nlines = append(nlines, scanner.Text())
		if len(nlines) > b {
			nlines = nlines[1:]
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return foundStrings
}

func getLinesAfter(s *bufio.Scanner, lc int) []string {
	out := make([]string, 0, lc)

	for range lc {
		if !s.Scan() {
			break
        }
		out = append(out, s.Text())
	}

	return out
}

func resetReader(r io.Reader) error {
	if seeker, ok := r.(io.Seeker); ok {
		_, err := seeker.Seek(0, io.SeekStart)
		return err
	}

	return nil
}

func WriteToFile(text, outputFile string) error {
	file, err := os.Create(outputFile)
	if err != nil {
		return fmt.Errorf("failed to create/open output file")
	}

	_, err = file.WriteString(text)
	if err != nil {
		return fmt.Errorf("failed to write output to %s", outputFile)
	}

	return nil
}
