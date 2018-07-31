package main

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"
)

/*
 * ホームディレクトリのパスを取得します
 */
func getHomeDir() string {
	me, err := user.Current()
	if err != nil {
		panic(err)
	}
	return me.HomeDir
}

// ホームディレクトリ(~)を展開するGolang APIはない(OSではなくシェルの機能のため)
func clean2(path string) string {
	if len(path) > 1 && path[0:2] == "~"+string(filepath.Separator) {
		path = getHomeDir() + path[1:]
	}
	path = os.ExpandEnv(path)
	return filepath.Clean(path)
}

func main() {
	fmt.Println(clean2("~\\file.txt"))
}
