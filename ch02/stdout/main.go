package main

import (
	"os"
)

func main() {
	// fmt.Printf内ではos.Stdout.Writeを呼び出しているので以下のコードは等価
	os.Stdout.Write([]byte("os.Stdout example\n"))
	// == fmt.Printf("os.Stdout example\n")
}
