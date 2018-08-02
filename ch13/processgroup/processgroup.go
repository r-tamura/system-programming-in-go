package main

import (
	"fmt"
	"os"
	"syscall"
)

// Only for Linux, MacOS
func main() {
	sid, _ := syscall.Getsid(os.Getpid())
	fmt.Printf("Group ID: %d\nSession Group ID: %d\n", syscall.Getpgrp(), sid)
}
