package main

import (
	"io"
	"os"
	"strings"
)

var (
	computer    = strings.NewReader("COMPUTER")
	system      = strings.NewReader("SYSTEM")
	programming = strings.NewReader("PROGRAMMING")
)

func main() {
	var stream io.Reader

	a := io.NewSectionReader(programming, 5, 1)
	s := io.NewSectionReader(system, 0, 1)
	c := io.NewSectionReader(computer, 0, 1)
	i := io.NewSectionReader(programming, 8, 1)

	pipeReader, pipeWriter := io.Pipe()
	writer := io.MultiWriter(pipeWriter, pipeWriter)
	go io.Copy(writer, i)
	defer pipeWriter.Close()

	stream = io.MultiReader(a, s, c, io.LimitReader(pipeReader, 2))

	io.Copy(os.Stdout, stream)
}
