package main

import (
	"fmt"
	"sync"
	"time"
)

func reader(id int, ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		val, ok := <-ch
		if !ok {
			fmt.Println("Channel Closed")
			return
		}
		var res int64 = 1
		for i := 0; i < 1000000000; i++ {
			res += 1
		}
		fmt.Printf("Reader: %d, Recieved: %d\n", id, val)
	}
}

func main() {
	wg := sync.WaitGroup{}
	ch := make(chan int)

	const N = 6

	wg.Add(N)

	for i := 0; i < N; i++ {
		go reader(i, ch, &wg)
	}

	t0 := time.Now()

	for i := 0; i < 100; i++ {
		ch <- i
	}

	close(ch)

	wg.Wait()

	fmt.Printf("Elapsed: %v\n", time.Since(t0))
}
