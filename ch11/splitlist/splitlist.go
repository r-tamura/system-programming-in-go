package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func getArgs(args []string) []string {
	if len(args) == 1 {
		fmt.Printf("%s [exec file name]", args[0])
		os.Exit(1)
	}

	firstArg := args[1]
	if strings.Index(firstArg, "--") != -1 {
		return args[2:]
	} else {
		return args[1:]
	}
}

func main() {
	args := getArgs(os.Args)
	for _, path := range filepath.SplitList(os.Getenv("PATH")) {
		execpath := filepath.Join(path, args[0])
		_, err := os.Stat(execpath)
		if !os.IsNotExist(err) {
			fmt.Println(execpath)
			return
		}
	}
	os.Exit(1)
}
