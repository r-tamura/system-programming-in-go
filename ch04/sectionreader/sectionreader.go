package main

import (
	"io"
	"os"
	"strings"
)

func main() {
	reader := strings.NewReader("A newreader\nAnother reader\n")
	sectionReader := io.NewSectionReader(reader, 14, 7)
	io.Copy(os.Stdout, sectionReader)
}
