package main

import (
	"archive/zip"
	"io"
	"net/http"
	"os"
	"strings"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/zip")
	w.Header().Set("Content-Disposition",
		"attachment; filename=sample.zip")

	file, err := os.Create("sample.zip")
	if err != nil {
		panic(err)
	}
	// Writer => zip.Writerを生成
	zipWriter := zip.NewWriter(file)
	defer zipWriter.Close()

	a, err := zipWriter.Create("a.txt")
	if err != nil {
		panic(err)
	}
	io.Copy(a, strings.NewReader("This is joke"))
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
