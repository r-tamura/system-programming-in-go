package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan bool, 1)
	fmt.Println("All tasks are started")
	go func() {
		time.Sleep(3 * time.Second)
		fmt.Println("Task is finished")
		done <- true
	}()

	<-done
	fmt.Println("All tasks are finished")
}
