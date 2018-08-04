package main

import (
	"fmt"
	"syscall"
)

/* 全てのOSで使用できることが保証されているシグナル */
const (
	interrupt syscall.Signal = syscall.SIGINT
	kill      syscall.Signal = syscall.SIGKILL
)

func main() {
	fmt.Printf("SIGINT: %d\n", interrupt)
	fmt.Printf("SIGKILL: %d\n", kill)
}
