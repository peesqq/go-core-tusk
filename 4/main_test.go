package main

import (
	"reflect"
	"testing"
)

func TestChangeSlice(t *testing.T) {
	slice1 := []string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"}
	slice2 := []string{"banana", "date", "fig"}
	res := []string{"banana", "date"}
	if !reflect.DeepEqual(ChangeSlice(slice1, slice2), res) {
		t.Errorf("Expected slice \"banana\", \"date\", but got other")
	}
}
