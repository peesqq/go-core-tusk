package main

import (
	"testing"
	"time"
)

func TestSemaphore(t *testing.T) {
	sem := NewSemaphore(0)
	var tasks []int

	for i := 1; i <= 3; i++ {
		sem.Add(1)
		go func(i int) {
			defer sem.Done()
			time.Sleep(100 * time.Millisecond)
			tasks = append(tasks, i)
		}(i)
	}

	sem.Wait()

	if len(tasks) != 3 {
		t.Errorf("Expected 3 tasks, got %d", len(tasks))
	}
}
