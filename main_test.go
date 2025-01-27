package main

import (
	"bytes"
	"os"
	"strings"
	"testing"
)

func TestMain_NoArgs(t *testing.T) {
	originalStderr := os.Stderr
	defer func() { os.Stderr = originalStderr }()

	r, w, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}

	os.Stderr = w
	os.Args = []string{"csv-split"}

	main()

	w.Close()

	var buf bytes.Buffer
	if _, err := buf.ReadFrom(r); err != nil {
		t.Fatal(err)
	}

	// 標準エラー出力の内容をトリム
	stderr := strings.TrimSpace(buf.String())
	expected := "Usage: csv-split <input-file> <lines-per-file>"
	if stderr != expected {
		t.Errorf("Unexpected error message: %q, expected: %q", stderr, expected)
	}
}
