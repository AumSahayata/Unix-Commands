package internal

import (
	"bufio"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

func OpenFile(filepath string) (*os.File, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func CountLines(rd io.Reader) (int, error) {
	count := 0
    if err := resetReader(rd); err != nil {
        return 0, err
    }

	scanner := bufio.NewScanner(rd)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan(){
		count++
	}

	if err := scanner.Err(); err != nil {
		return 0, err
	}

	return count, nil
}

func CountWords(rd io.Reader) (int, error) {
	count := 0
    if err := resetReader(rd); err != nil {
        return 0, err
    }

	scanner := bufio.NewScanner(rd)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan(){
		token := scanner.Text()
		if isWordToken(token) {
			count++
		}
	}

	if err := scanner.Err(); err != nil {
		return 0, err
	}

	return count, nil
}

func isEmoji(r rune) bool {
	// Filter common emoji Unicode ranges
	return (r >= 0x1F600 && r <= 0x1F64F) || // emoticons
		(r >= 0x1F300 && r <= 0x1F5FF) || // symbols & pictographs
        (r >= 0x1F680 && r <= 0x1F6FF) || // transport
        (r >= 0x1F1E6 && r <= 0x1F1FF) // flags
}

// using this because a token can be 'hello😐' and that needs to passed one by one
func isWordToken(token string) bool {
	for _, r := range token {
		if isEmoji(r) {
			return false
		}
	}
	return true
}

func CountCharacters(rd io.Reader) (int, error) {
	count := 0
    if err := resetReader(rd); err != nil {
        return 0, err
    }
	
	scanner := bufio.NewScanner(rd)
	scanner.Split(bufio.ScanRunes)

	for scanner.Scan(){
		r, _ := utf8.DecodeRune(scanner.Bytes())
		if !unicode.IsSpace(r) && !isEmoji(r) && r != '\n' {
			count++
		}
	}

	if err := scanner.Err(); err != nil {
		return 0, err
	}

	return count, nil
}

func resetReader(r io.Reader) error {
	if seeker, ok := r.(io.Seeker); ok {
		_, err := seeker.Seek(0, io.SeekStart)
		return err
	}

	return nil
}