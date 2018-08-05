package main

import (
	"fmt"
	"sync"
	"time"
)

// チャネルの場合、ブロックしているgoroutineに通知する場合チャネルをクローズするしかない
// sync.Condは何度でも使える
func main() {
	var mtx sync.Mutex
	cond := sync.NewCond(&mtx)

	for _, name := range []string{"A", "B", "C"} {
		go func(name string) {
			mtx.Lock()
			defer mtx.Unlock()
			cond.Wait()
			fmt.Println(name)
		}(name)
	}
	// 3つのgoroutineがcond.Wait()されるの待つ
	time.Sleep(time.Second)
	// waitしているgoroutineを起こす
	cond.Broadcast()
	// main関数が終了してしまうとプログラムが終了してしまうのでまつ
	time.Sleep(time.Second)
}
