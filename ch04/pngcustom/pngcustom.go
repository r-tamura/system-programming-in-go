package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"hash/crc32"
	"io"
	"os"
)

func dumpChunk(chunk io.Reader) {
	var length int32
	binary.Read(chunk, binary.BigEndian, &length)
	buffer := make([]byte, 4)
	chunk.Read(buffer)
	fmt.Printf("chunk '%v' (%d bytes)\n", string(buffer), length)
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

func textChunk(text string) io.Reader {
	byteData := []byte(text)
	var buffer bytes.Buffer

	binary.Write(&buffer, binary.BigEndian, int32(len(byteData)))
	buffer.WriteString("tEXt")
	buffer.Write(byteData)
	crc := crc32.NewIEEE()
	io.WriteString(crc, "tEXt")
	binary.Write(&buffer, binary.BigEndian, crc.Sum32())
	return &buffer
}

func main() {
	png, err := os.Open("Lenna.png")
	if err != nil {
		panic(err)
	}
	defer png.Close()

	newPng, err := os.Create("new_Lenna.png")
	if err != nil {
		panic(err)
	}
	defer newPng.Close()

	chunks := readChunks(png)

	// Add signiture
	io.WriteString(newPng, "\x89PNG\r\n\x1a\n")
	// Add IHDR chunk
	io.Copy(newPng, chunks[0])

	// Add Custom chunk
	io.Copy(newPng, textChunk("ASCII PROGRAMMING++"))

	// Add rest chunks
	for _, chunk := range chunks[1:] {
		io.Copy(newPng, chunk)
	}
}
