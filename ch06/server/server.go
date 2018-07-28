package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/http/httputil"
	"strings"
)

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
		go func() {
			fmt.Printf("Accept %v\n", connection.RemoteAddr())

			// Read the request
			request, err := http.ReadRequest(bufio.NewReader(connection))
			if err != nil {
				panic(err)
			}

			// デバッグ 第二引数をtrueにすることでbodyもダンプする
			dump, err := httputil.DumpRequest(request, true)
			if err != nil {
				panic(err)
			}
			fmt.Println(string(dump))
			response := http.Response{
				StatusCode: 200,
				ProtoMajor: 1,
				ProtoMinor: 0,
				Body:       ioutil.NopCloser(strings.NewReader("Hello, Golang\n")),
			}
			response.Write(connection)
			connection.Close()
		}()
	}
}
