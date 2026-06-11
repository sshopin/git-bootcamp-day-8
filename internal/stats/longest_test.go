package stats

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLongestLine(t *testing.T) {
	path := filepath.Join(t.TempDir(), "sample.txt")
	content := "short\nlonger line here\nmid\n"
	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		t.Fatal(err)
	}

	got, err := LongestLine(path)
	if err != nil {
		t.Fatal(err)
	}
	if got != 16 {
		t.Fatalf("LongestLine() = %d, want 16", got)
	}
}

func TestLongestLineUnicode(t *testing.T) {
	path := filepath.Join(t.TempDir(), "sample.txt")
	content := "café\nширокая строка\n"
	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		t.Fatal(err)
	}

	got, err := LongestLine(path)
	if err != nil {
		t.Fatal(err)
	}
	if got != 14 {
		t.Fatalf("LongestLine() = %d, want 14", got)
	}
}
