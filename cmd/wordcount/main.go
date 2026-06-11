package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/git-bootcamp/wordcount/internal/stats"
)

func main() {
	showLines := flag.Bool("lines", false, "count lines")
	showWords := flag.Bool("words", false, "count words")
	showBytes := flag.Bool("bytes", false, "count bytes")
	showChars := flag.Bool("chars", false, "count Unicode characters")
	showLongest := flag.Bool("longest", false, "length of the longest line in Unicode characters")
	flag.Parse()

	if flag.NArg() != 1 {
		fmt.Fprintln(os.Stderr, "usage: wordcount [--lines|--words|--bytes|--chars|--longest] <file>")
		os.Exit(2)
	}

	path := flag.Arg(0)
	selected := 0
	if *showLines {
		selected++
	}
	if *showWords {
		selected++
	}
	if *showBytes {
		selected++
	}
	if *showChars {
		selected++
	}
	if *showLongest {
		selected++
	}
	if selected != 1 {
		fmt.Fprintln(os.Stderr, "error: specify exactly one of --lines, --words, --bytes, --chars, --longest")
		os.Exit(2)
	}

	var (
		value int
		err   error
	)
	switch {
	case *showLines:
		value, err = stats.CountLines(path)
	case *showWords:
		value, err = stats.CountWords(path)
	case *showBytes:
		value, err = stats.CountBytes(path)
	case *showChars:
		value, err = stats.CountChars(path)
	case *showLongest:
		value, err = stats.LongestLine(path)
	}
	if err != nil {
		fmt.Fprintf(os.Stderr, "wordcount: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(value)
}
