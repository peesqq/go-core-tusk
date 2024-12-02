package main

import (
	"testing"
)

func TestGenerateRandomSlice(t *testing.T) {
	slice := generateRandomSlice(10)
	if len(slice) != 10 {
		t.Errorf("Expected slice length 10, but got %d", len(slice))
	}
}

func TestSliceExample(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5}
	expected := []int{2, 4}
	result := sliceExample(slice)
	for i, v := range expected {
		if result[i] != v {
			t.Errorf("Expected %d at index %d, but got %d", v, i, result[i])
		}
	}
}

func TestAddElements(t *testing.T) {
	slice := []int{1, 2, 3}
	elem := 4
	expected := []int{1, 2, 3, 4}
	result := addElements(slice, elem)
	for i, v := range expected {
		if result[i] != v {
			t.Errorf("Expected %d at index %d, but got %d", v, i, result[i])
		}
	}
}

func TestCopySlice(t *testing.T) {
	slice := []int{1, 2, 3}
	copiedSlice := copySlice(slice)
	slice[0] = 99 // Изменяем оригинал

	if copiedSlice[0] == 99 {
		t.Errorf("Expected copied slice to remain unchanged, but got %v", copiedSlice)
	}
}

func TestRemoveElement(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5}
	index := 2
	expected := []int{1, 2, 4, 5}
	result := removeElement(slice, index)
	for i, v := range expected {
		if result[i] != v {
			t.Errorf("Expected %d at index %d, but got %d", v, i, result[i])
		}
	}
}
