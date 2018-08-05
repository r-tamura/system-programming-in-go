package main

import (
	"fmt"
	"sync"
)

func initialize() {
	fmt.Println("Intilization")
}

// initという関数名があるとその関数が一度だけ初期化処理として呼び出されるのでsync.Onceよりそちらのが良い
// 初期化処理を遅延させる場合にsync.Onceを使う
func main() {
	var once sync.Once
	once.Do(initialize)
	once.Do(initialize)
	once.Do(initialize)
}
