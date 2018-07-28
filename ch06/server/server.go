package main

import (
	"bufio"
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"net/http/httputil"
	"strings"
	"time"
)

func isGZipAcceptable(request *http.Request) bool {
	return strings.Index(strings.Join(request.Header["Accept-Encoding"], ","), "gzip") != -1
}

func processSession(connection net.Conn) {
	fmt.Printf("Accept %v\n", connection.RemoteAddr())

	// To loop for responding all time after acceptance
	for {
		// Setup time out settings
		connection.SetReadDeadline(time.Now().Add(5 * time.Second))

		// Read the request
		request, err := http.ReadRequest(bufio.NewReader(connection))
		if err != nil {
			// When time out or socket is closed, this program will exit,
			// otherwise throws an error.
			neterr, ok := err.(net.Error) // down cast
			if ok && neterr.Timeout() {
				fmt.Println("Timeout")
				break
			} else if err == io.EOF {
				break
			}
			panic(err)
		}

		// デバッグ 第二引数をtrueにすることでbodyもダンプする
		dump, err := httputil.DumpRequest(request, true)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(dump))

		content := "Hello, Golang\n"
		response := http.Response{
			StatusCode:    200,
			ProtoMajor:    1,
			ProtoMinor:    1,
			ContentLength: int64(len(content)),
			Header:        make(http.Header),
		}

		// GZip or raw
		if isGZipAcceptable(request) {
			content := "<h1>Hello, Golang<span>(gzipped)</span></h1>"
			// Respond with the gzipped content
			var buffer bytes.Buffer
			writer := gzip.NewWriter(&buffer)
			io.WriteString(writer, content)
			writer.Close()
			response.Body = ioutil.NopCloser(&buffer)
			response.ContentLength = int64(buffer.Len())
			response.Header.Set("Content-Encoding", "gzip")
		} else {
			content := "<h1>Hello, Golang</h1>\n"
			response.Body = ioutil.NopCloser(strings.NewReader(content))
		}

		response.Write(connection)
	}
	connection.Close()
}

var (
	host = "localhost"
	port = 8080
)

func main() {
	listen := fmt.Sprintf("%v:%v", host, port)
	listener, err := net.Listen("tcp", listen)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Server is running at %v\n", listen)

	for {
		connection, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		go processSession(connection)
	}
}
