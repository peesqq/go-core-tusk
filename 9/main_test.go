package main

import (
	"testing"
)

func TestNumberPipeline(t *testing.T) {
	input := make(chan uint8)
	output := make(chan float64)

	go NumberPipeline(input, output)

	go func() {
		for _, num := range []uint8{1, 2, 3, 4} {
			input <- num
		}
		close(input)
	}()

	expected := []float64{1, 8, 27, 64}

	i := 0
	for result := range output {
		if result != expected[i] {
			t.Errorf("Expected %.2f, got %.2f", expected[i], result)
		}
		i++
	}

	if i != len(expected) {
		t.Errorf("Expected %d results, got %d", len(expected), i)
	}
}
