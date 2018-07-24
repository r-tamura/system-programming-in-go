package main

import (
	"archive/zip"
	"io"
	"os"
	"strings"
)

func addFile(zipWriter *zip.Writer, filename string, str string) {
	writer, err := zipWriter.Create(filename)
	if err != nil {
		panic(err)
	}
	io.Copy(writer, strings.NewReader(str))
}

func main() {
	file, err := os.Create("result.zip")
	if err != nil {
		panic(err)
	}

	zipWriter := zip.NewWriter(file)
	defer zipWriter.Close()

	addFile(zipWriter, "a.txt", "テキストa")
	addFile(zipWriter, "b.txt", "テキストb")
}
