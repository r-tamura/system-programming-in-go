package main

import (
	"bufio"
	"compress/gzip"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"strings"
)

func main() {

	sendMessages := []string{
		"ASCII", "PROGRAMMING", "PLUS",
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

		request.Write(connection)
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
		defer response.Body.Close()

		if response.Header.Get("Content-Encoding") == "gzip" {
			reader, err := gzip.NewReader(response.Body)
			if err != nil {
				panic(err)
			}
			io.Copy(os.Stdout, reader)
		} else {
			io.Copy(os.Stdout, response.Body)
		}

		// If all request is sent, exit
		current++
		if current == len(sendMessages) {
			break
		}
	}
}
