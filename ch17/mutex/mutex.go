package main

import (
	"fmt"
	"sync"
)

var id int

func generateId(mutex *sync.Mutex, wg *sync.WaitGroup) int {
	// Lock()/Unlock()
	mutex.Lock()
	defer mutex.Unlock()
	defer wg.Done()
	id++
	return id
}

func main() {
	var wg sync.WaitGroup
	var mutex sync.Mutex

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			fmt.Printf("id: %d\n", generateId(&mutex, &wg))
		}()
	}
	wg.Wait()
	fmt.Println("End")
}
