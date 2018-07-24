package main

import (
	"io"
	"os"
)

func main() {
	from := os.Args[1]
	to := os.Args[2]

	source, err := os.Open(from)
	defer source.Close()
	if err != nil {
		panic(err)
	}
	dest, err := os.Create(to)
	defer dest.Close()
	if err != nil {
		panic(err)
	}
	io.Copy(dest, source)
}
