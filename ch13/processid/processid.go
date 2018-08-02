package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("Process ID: %d\n", os.Getpid())
	fmt.Printf("Parent Process ID: %d\n", os.Getppid())
}
