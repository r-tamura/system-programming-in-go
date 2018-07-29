package main

import (
	"bufio"
	"fmt"
	"net"
	"net/http"
	"net/http/httputil"
)

func main() {

	sendMessages := []string{
		"ASCII", "PROGRAMMING", "PLUS",
	}

	current := 0
	var connection net.Conn
	var err error
	requests := make(chan *http.Request, len(sendMessages))
	// when a connection is not created yet or retrying for error, start with dial up.
	if connection == nil {
		connection, err = net.Dial("tcp", "localhost:8080")
		if err != nil {
			panic(err)
		}
		fmt.Printf("Access: %d\n", current)
	}

	// A loop for retry
	for i := 0; i < len(sendMessages); i++ {
		// 最後のリクエストであるか
		lastMessage := i == len(sendMessages)-1
		request, err := http.NewRequest("GET", "http://localhost:8080?message="+sendMessages[i], nil)

		// 最後のリクエストの場合はコネクションを閉じる
		if lastMessage {
			request.Header.Add("Connection", "close")
		} else {
			request.Header.Add("Connection", "keep-alive")
		}

		if err != nil {
			panic(err)
		}

		err = request.Write(connection)
		if err != nil {
			panic(err)
		}
		fmt.Println("send: ", sendMessages[i])
		requests <- request
	}

	close(requests)

	// Recieve all response at once
	reader := bufio.NewReader(connection)

	for request := range requests {
		response, err := http.ReadResponse(reader, request)
		if err != nil {
			panic(err)
		}
		dump, err := httputil.DumpResponse(response, true)
		if err != nil {
			panic(err)
		}

		fmt.Println(string(dump))
		current++
		if current == len(sendMessages) {
			break
		}
	}

}
