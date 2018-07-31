package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	fmt.Println(filepath.Clean("./filepath//../path.go")) // path.go
	fmt.Println(filepath.Abs("./clean/path.go"))          // D:\Git\GitHub\system-programming-in-go\ch11\clean\path.go
}
