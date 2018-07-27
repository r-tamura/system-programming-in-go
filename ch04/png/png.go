package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"os"
)

func dumpChunk(chunk io.Reader) {
	var length int32
	binary.Read(chunk, binary.BigEndian, &length)
	buffer := make([]byte, 4)
	chunk.Read(buffer)

	fmt.Printf("chunk '%v' (%d bytes)\n", string(buffer), length)
	if bytes.Equal(buffer, []byte("tEXt")) {
		rawText := make([]byte, length)
		chunk.Read(rawText)
		fmt.Println(string(rawText))
	}
}

func readChunks(file *os.File) []io.Reader {

	var chunks []io.Reader

	// Skip first 8 byte
	file.Seek(8, 0)
	var offset int64 = 8

	for {
		var length int32
		err := binary.Read(file, binary.BigEndian, &length)
		if err == io.EOF {
			break
		}
		chunks = append(chunks, io.NewSectionReader(file, offset, int64(length)+12))

		// Move to first byte of next chunk
		// チャンク名(4バイト) + データ長 + CRC(4バイト)先に移動
		offset, _ = file.Seek(int64(4+length+4), 1)
	}

	return chunks
}

func main() {
	png, err := os.Open("new_Lenna.png")
	if err != nil {
		panic(err)
	}

	chunks := readChunks(png)
	for _, chunk := range chunks {
		dumpChunk(chunk)
	}
}
