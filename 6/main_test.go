package main

import (
	"math/rand"
	"testing"
	"time"
)

func TestRandomNumberGenerator(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	ch := make(chan int)

	go RandomNumberGenerator(ch, 5)

	expectedCount := 5
	count := 0
	for range ch {
		count++
	}

	if count != expectedCount {
		t.Errorf("Expected %d numbers, but got %d", expectedCount, count)
	}
}
