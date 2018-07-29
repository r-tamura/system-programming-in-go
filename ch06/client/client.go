package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/http/httputil"
	"strconv"
	"strings"
)

func main() {

	sendMessages := []string{
		"ASCII",
	}

	current := 0
	var connection net.Conn
	// A loop for retry
	for {
		var err error
		// when a connection is not created yet or retrying for error, start with dial up.
		if connection == nil {
			connection, err = net.Dial("tcp", "localhost:8080")
			if err != nil {
				panic(err)
			}
			fmt.Printf("Access: %d\n", current)
		}

		// Fetch from the server. If timeout has caused, retry, because an error is emmited in this line.
		request, err := http.NewRequest("POST", "http://localhost:8080", strings.NewReader(sendMessages[current]))
		if err != nil {
			panic(err)
		}
		request.Header.Set("Accept-Encoding", "gzip")

		err = request.Write(connection)
		if err != nil {
			panic(err)
		}
		reader := bufio.NewReader(connection)

		response, err := http.ReadResponse(bufio.NewReader(connection), request)
		if err != nil {
			fmt.Println("Retry")
			connection = nil
			continue
		}

		dump, err := httputil.DumpResponse(response, false)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(dump))

		// Chunked response
		if len(response.TransferEncoding) < 1 || response.TransferEncoding[0] != "chunked" {
			panic("Wrng transfer encoding")
		}

		for {
			// Get the size of a chunk
			sizeStr, err := reader.ReadBytes([]byte("\n")[0])
			if err == io.EOF {
				break
			}
			// Parse the size in hex. If the size is 0, close.
			size, err := strconv.ParseInt(string(sizeStr[:len(sizeStr)-2]), 16, 64)
			if size == 0 {
				break
			}
			if err != nil {
				panic(err)
			}

			// Allocate memory only the size
			line := make([]byte, int(size))
			reader.Read(line)
			// 2続く改行コードの2つ目をスキップする
			reader.Discard(2)
			fmt.Printf("  %d bytes: %s\n", size, string(line))
		}

		// If all request is sent, exit
		current++
		if current == len(sendMessages) {
			break
		}
	}
}
