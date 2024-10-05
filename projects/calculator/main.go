package main

import (
	"fmt"
	"sync"
)

func calculator(firstChan <-chan int, secondChan <-chan int, stopChan <-chan struct{}) <-chan int {
	ch := make(chan int)
	go func() {
		select {
		case inp := <-firstChan:
			ch <- inp * inp

		case inp := <-secondChan:
			ch <- inp * 3

		case _ = <-stopChan:
			break
		}
		close(ch)
	}()
	return ch
}

func main() {
	fch := make(chan int)
	sch := make(chan int)
	stop := make(chan struct{})
	ch := calculator(fch, sch, stop)

	wg := sync.WaitGroup{}

	wg.Add(1)

	go func() {
		for ans := range ch {
			fmt.Println(ans)
		}
		wg.Done()
	}()

	fch <- 10

	wg.Wait()
}
