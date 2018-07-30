package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Printf("%s [exec file name]", os.Args[0])
		os.Exit(1)
	}

	targetFile := os.Args[1]
	if strings.Index(targetFile, "-") != -1 {
		targetFile = os.Args[2]
	}
	info, err := os.Stat(targetFile)
	if err == os.ErrNotExist {
		fmt.Printf("file not found: %s", targetFile)
	} else if err != nil {
		panic(err)
	}

	fmt.Println("FileInfo")
	fmt.Printf("  File name: %v\n", info.Name())
	fmt.Printf("  Size: %v\n", info.Size())
	fmt.Printf("  Last modified: %v\n", info.ModTime())
	fmt.Println("Mode()")
	fmt.Printf("  directory?: %v\n", info.Mode().IsDir())
	fmt.Printf("  readable and writable?: %v\n", info.Mode().IsRegular())
}
