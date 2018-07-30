package main

import (
	"io"
	"os"
)

func append(path string) {
	file, err := os.OpenFile(path, os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	io.WriteString(file, "Append content\n")
}

func main() {
	append("append\\text.txt")
}
