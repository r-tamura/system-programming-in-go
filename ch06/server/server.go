package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"net/http/httputil"
	"strings"
	"time"
)

// 順番に従ってconnnectionへwriteする(this func should be invoked as go routine)
func writeToConn(sessionResponses chan chan *http.Response, connection net.Conn) {
	defer connection.Close()
	// Take responses in turn
	for sessionResponse := range sessionResponses {
		// Wait for the selected work
		response := <-sessionResponse
		response.Write(connection)
		close(sessionResponse)
	}
}

func handleRequest(request *http.Request, resultReceiver chan *http.Response) {
	// デバッグ 第二引数をtrueにすることでbodyもダンプする
	dump, err := httputil.DumpRequest(request, true)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(dump))

	content := "<h1>Hello, Golang</h1>\n"
	response := &http.Response{
		StatusCode:    200,
		ProtoMajor:    1,
		ProtoMinor:    1,
		ContentLength: int64(len(content)),
		Body:          ioutil.NopCloser(strings.NewReader(content)),
	}

	// Write to a channel ffter a process finished and
	// restart the blocked writeToConn process
	resultReceiver <- response
}

func processSession(connection net.Conn) {
	fmt.Printf("Accept %v\n", connection.RemoteAddr())

	// Channel for processing requests in right order in an session
	sessionResponses := make(chan chan *http.Response, 50)
	defer close(sessionResponses)

	// Goroutine to serialize responses and write to socket.
	go writeToConn(sessionResponses, connection)

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
		sessionResponse := make(chan *http.Response)
		sessionResponses <- sessionResponse
		// Run async response
		go handleRequest(request, sessionResponse)
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
