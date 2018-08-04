package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	fmt.Println("All tasks are started")
	// 終了を受け取るための終了関数付きコンテキスト
	// A context with the completing function for receiving the notification for completed.
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		time.Sleep(3 * time.Second)
		fmt.Println("Task is finished")
		// 終了を通知
		// Notify the task was completed
		cancel()
	}()

	// wait for fin(ish of the task
	<-ctx.Done()

	fmt.Println("All tasks are finished")

}
