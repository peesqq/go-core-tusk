package main

import (
	"reflect"
	"sort"
	"testing"
)

func TestMergeChannels(t *testing.T) {
	// Входные каналы
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)

	// Запуск горутин для передачи данных
	go func() {
		defer close(ch1)
		ch1 <- 1
		ch1 <- 2
	}()

	go func() {
		defer close(ch2)
		ch2 <- 3
		ch2 <- 4
	}()

	go func() {
		defer close(ch3)
		ch3 <- 5
	}()

	// Ожидаемый результат (порядок не важен)
	expected := []int{1, 2, 3, 4, 5}

	// Слияние каналов
	merged := MergeChannels([]<-chan int{ch1, ch2, ch3})

	// Чтение результата из объединённого канала
	var result []int
	for val := range merged {
		result = append(result, val)
	}
	sort.Ints(result)

	// Проверка результата
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected %v, got %v", expected, result)
	}
}
