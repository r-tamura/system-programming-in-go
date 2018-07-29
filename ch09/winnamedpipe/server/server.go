package main

// GoではWindows名前付きパイプを機能は提供されていない
// 以下のライブラリで使用することができる
import (
	"fmt"
	"net"

	"gopkg.in/natefinch/npipe.v2"
)

func handleConnection(conn net.Conn) {
	buffer := make([]byte, 500)
	length, err := conn.Read(buffer)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(buffer[:length]))
	conn.Write([]byte("Hello from Windows named pipe sever!\n"))
	defer conn.Close()
}

func main() {
	fmt.Println("Listen on Windows named pipe")
	listener, err := npipe.Listen("\\\\.\\pipe\\mypipe")
	if err != nil {
		panic(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		go handleConnection(conn)
	}
}
