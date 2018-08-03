package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	args := getArgs(os.Args)

	if len(args) == 1 {
		return
	}

	cmd := exec.Command(args[1], args[2:]...)
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
	state := cmd.ProcessState
	fmt.Printf("%s\n", state.String())
	fmt.Printf("  isSuccess: %t\n", state.Success())
	fmt.Printf("  Pid: %d\n", state.Pid())
	fmt.Printf("  System: %v\n", state.SystemTime())
	fmt.Printf("  User: %v\n", state.UserTime())
}
