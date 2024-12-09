package main

import (
	"sync"
)

type Semaphore struct {
	count int
	ch    chan struct{}
	mu    sync.Mutex
}

func NewSemaphore(count int) *Semaphore {
	return &Semaphore{
		count: count,
		ch:    make(chan struct{}, count),
	}
}

func (s *Semaphore) Add(n int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.count += n
}

func (s *Semaphore) Done() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.count--
	if s.count == 0 {
		close(s.ch)
	}
}

func (s *Semaphore) Wait() {
	<-s.ch
}
