package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func getArgs(args []string) []string {
	if len(args) == 1 {
		fmt.Printf("expected 1 args, but got %v\n", len(args))
		os.Exit(1)
	}

	firstArg := args[1]
	if strings.Index(firstArg, "--") != -1 {
		return args[2:]
	} else {
		return args[1:]
	}
}

var jsSuffix = map[string]bool{
	".js":  true,
	".jsx": true,
	".tx":  true,
	".tsx": true,
}

func main() {
	args := getArgs(os.Args)
	root := args[0]
	err := filepath.Walk(root,
		func(path string, info os.FileInfo, err error) error {

			if info.IsDir() {
				if info.Name() == "node_modules" {
					return filepath.SkipDir
				}
				return nil
			}

			ext := strings.ToLower(filepath.Ext(info.Name()))
			if jsSuffix[ext] {
				rel, err := filepath.Rel(root, path)
				if err != nil {
					return nil
				}
				fmt.Printf("%s\n", rel)
			}
			return nil
		})
	if err != nil {
		fmt.Println(1, err)
	}
}
