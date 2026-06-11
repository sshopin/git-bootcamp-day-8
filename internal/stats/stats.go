package stats

import (
	"bufio"
	"os"
	"strings"
	"unicode/utf8"
)

// CountLines returns the number of lines in a text file.
func CountLines(path string) (int, error) {
	f, err := os.Open(path)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	n := 0
	for scanner.Scan() {
		n++
	}
	return n, scanner.Err()
}

// CountWords returns the number of whitespace-separated words.
func CountWords(path string) (int, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return 0, err
	}
	return len(strings.Fields(string(data))), nil
}

// CountBytes returns the file size in bytes.
func CountBytes(path string) (int, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return 0, err
	}
	return len(data), nil
}

// CountChars returns the number of Unicode code points.
func CountChars(path string) (int, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return 0, err
	}
	return utf8.RuneCount(data), nil
}
