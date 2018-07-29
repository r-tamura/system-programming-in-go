package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/http/httputil"
	"strings"
	"time"
)

var contents = []string{
	"Lorem Ipsum is simply dummy text of the printing and typesetting industry.",
	"Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, ",
	"when an unknown printer took a galley of type and scrambled it to make a type specimen book.",
	"It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged.",
	"It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, ",
	"and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.",
}

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

		fmt.Fprint(connection, strings.Join([]string{
			"HTTP/1.1 200 OK",
			"Centent-Type: text/plain",
			"Transfer-Encoding: chunked",
			"", "",
		}, "\r\n"))

		// Write response body
		for _, content := range contents {
			bytes := []byte(content)
			fmt.Fprintf(connection, "%x\r\n%s\r\n", len(bytes), content)
		}

		fmt.Fprintf(connection, "0\r\n\r\n")
	}
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
