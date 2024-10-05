package main

import (
	"fmt"
	"sync"
)

func removeDuplicates(inputStream chan string, outputStream chan string) {
	var last string
	for s := range inputStream {
		if s != last {
			outputStream <- s
		}
		last = s
	}
	close(outputStream)
}

func main() {
	input := make(chan string)
	output := make(chan string)

	go removeDuplicates(input, output)

	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() {
		for ans := range output {
			fmt.Println(ans)
		}
		wg.Done()
	}()

	input <- "Hello"
	input <- "Hello"
	input <- "World"

	close(input)

	wg.Wait()
}
