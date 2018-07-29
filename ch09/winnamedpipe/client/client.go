package main

import (
	"bufio"
	"fmt"

	"gopkg.in/natefinch/npipe.v2"
)

func main() {
	conn, err := npipe.Dial("\\\\.\\pipe\\mypipe")
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(conn, "Hi server!\n")
	msg, err := bufio.NewReader(conn).ReadString([]byte("\n")[0])
	fmt.Println(msg)
}
