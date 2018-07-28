package main

import (
	"io"
	"os"
	"strings"
)

func copyN(sink io.Writer, src io.Reader, length int) {
	limitedSrc := io.LimitReader(src, int64(length))
	io.Copy(sink, limitedSrc)
}

func main() {
	reader := strings.NewReader("Sample text 30.")
	copyN(os.Stdout, reader, 3)
}
