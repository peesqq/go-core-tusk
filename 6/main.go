package main

import (
	"fmt"
	"math/rand"
	"time"
)

func RandomNumberGenerator(ch chan<- int, count int) {
	for i := 0; i < count; i++ {
		ch <- rand.Intn(100)
	}
	close(ch)
}

func main() {
	rand.Seed(time.Now().UnixNano())
	ch := make(chan int)
	go RandomNumberGenerator(ch, 5)

	for num := range ch {
		fmt.Println(num)
	}
}
