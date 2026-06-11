package stats

import (
	"os"
	"path/filepath"
	"testing"
)

func writeTempFile(t *testing.T, content string) string {
	t.Helper()
	dir := t.TempDir()
	path := filepath.Join(dir, "sample.txt")
	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		t.Fatal(err)
	}
	return path
}

func TestCountLines(t *testing.T) {
	path := writeTempFile(t, "one\ntwo\nthree\n")
	got, err := CountLines(path)
	if err != nil {
		t.Fatal(err)
	}
	if got != 3 {
		t.Fatalf("CountLines() = %d, want 3", got)
	}
}

func TestCountWords(t *testing.T) {
	path := writeTempFile(t, "hello world\nfrom wordcount\n")
	got, err := CountWords(path)
	if err != nil {
		t.Fatal(err)
	}
	if got != 4 {
		t.Fatalf("CountWords() = %d, want 4", got)
	}
}

func TestCountBytes(t *testing.T) {
	path := writeTempFile(t, "abc")
	got, err := CountBytes(path)
	if err != nil {
		t.Fatal(err)
	}
	if got != 3 {
		t.Fatalf("CountBytes() = %d, want 3", got)
	}
}

func TestCountCharsUnicode(t *testing.T) {
	path := writeTempFile(t, "café")
	got, err := CountChars(path)
	if err != nil {
		t.Fatal(err)
	}
	if got != 4 {
		t.Fatalf("CountChars() = %d, want 4", got)
	}
}
