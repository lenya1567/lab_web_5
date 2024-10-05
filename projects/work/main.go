package main

import (
	"fmt"
	"sync"
)

func work() {
	fmt.Println("Hello, World!")
}

func main() {
	wg := sync.WaitGroup{}
	wg.Add(10)

	for i := 0; i < 10; i++ {
		go func() {
			work()
			wg.Done()
		}()
	}

	wg.Wait()
}
