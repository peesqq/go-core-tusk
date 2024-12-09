package main

import (
	"fmt"
	"math"
)

func NumberPipeline(input <-chan uint8, output chan<- float64) {
	for num := range input {
		cubed := math.Pow(float64(num), 3)
		output <- cubed
	}
	close(output)
}

func main() {
	input := make(chan uint8)
	output := make(chan float64)

	go NumberPipeline(input, output)

	go func() {
		for i := uint8(1); i <= 5; i++ {
			input <- i
		}
		close(input)
	}()

	for result := range output {
		fmt.Printf("Result: %.2f\n", result)
	}
}
