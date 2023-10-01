package internal

import (
	"bytes"
	"fmt"
	"io"
	"path/filepath"
	"strings"
)

func ShowContent(r io.Reader) string {
	var bf bytes.Buffer
	bf.ReadFrom(r)
	return bf.String()
}

func Dates(todoRoot string) ([]string, error) {
	files, err := filepath.Glob(filepath.Join(todoRoot, "*.md"))
	if err != nil {
		return nil, fmt.Errorf("error while searching %s: %w", todoRoot, err)
	}
	dates := make([]string, len(files))
	for i, file := range files {
		dates[i] = strings.TrimRight(filepath.Base(file), ".md")
	}
	return dates, nil
}
