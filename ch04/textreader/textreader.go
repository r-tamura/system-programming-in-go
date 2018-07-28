package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

/*
 * bufio.Reader
 * io.Readerまたはio.Writerをバッファ機能付きのReader,Writerとしてラップする
 */

var source = `Line1
Line2
Line3
Line4
`

func main() {
	reader := bufio.NewReader(strings.NewReader(source))
	for {
		line, err := reader.ReadString([]byte("\n")[0])
		if err == io.EOF {
			break
		}
		fmt.Printf("%#v\n", line)
	}
}
