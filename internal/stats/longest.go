package stats

import (
	"bufio"
	"os"
	"unicode/utf8"
)

// LongestLine returns the length of the longest line in Unicode code points.
func LongestLine(path string) (int, error) {
	f, err := os.Open(path)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	max := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if n := utf8.RuneCountInString(line); n > max {
			max = n
		}
	}
	return max, scanner.Err()
}
