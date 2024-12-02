package main

import (
	"fmt"
	"math/rand"
)

func generateRandomSlice(size int) []int {
	slice := make([]int, size)
	for i := range slice {
		slice[i] = rand.Intn(1000)
	}
	return slice
}

func sliceExample(slice []int) []int {
	var res []int
	for _, num := range slice {
		if num%2 == 0 {
			res = append(res, num)
		}
	}
	return res
}

func addElements(slice []int, elem int) []int {
	return append(slice, elem)
}

func copySlice(slice []int) []int {
	copiedSlice := make([]int, len(slice))
	copy(copiedSlice, slice)
	return copiedSlice
}

func removeElement(slice []int, index int) []int {
	if index < 0 || index >= len(slice) {
		fmt.Println("Index out of range")
		return slice
	}
	return append(slice[:index], slice[index+1:]...) //умно умно
}

func main() {
	originalSlice := generateRandomSlice(10)
	fmt.Println("Original Slice:", originalSlice)

	evenSlice := sliceExample(originalSlice)
	fmt.Println("Even Numbers Slice:", evenSlice)

	newSlice := addElements(originalSlice, 42)
	fmt.Println("Slice after Adding 42:", newSlice)

	copiedSlice := copySlice(originalSlice)
	fmt.Println("Copied Slice:", copiedSlice)

	// Проверка независимости
	originalSlice[0] = 999
	fmt.Println("Modified Original Slice:", originalSlice)
	fmt.Println("Unchanged Copied Slice:", copiedSlice)

	// 5. Удаление элемента
	updatedSlice := removeElement(originalSlice, 3)
	fmt.Println("Slice after Removing Element at Index 3:", updatedSlice)
}
