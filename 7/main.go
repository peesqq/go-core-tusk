package main

import (
	"sync"
)

func MergeChannels(channels []<-chan int) <-chan int {
	res := make(chan int)
	var wg sync.WaitGroup

	for _, ch := range channels {
		wg.Add(1)
		go func(c <-chan int) {
			defer wg.Done()
			for val := range c {
				res <- val
			}
		}(ch)
	}

	go func() {
		wg.Wait()
		close(res)
	}()

	return res
}
