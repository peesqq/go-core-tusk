package main

import (
	"reflect"
	"testing"
)

func TestFindIntersections(t *testing.T) {
	tests := []struct {
		name     string
		a        []int
		b        []int
		expected struct {
			found       bool
			intersected []int
		}
	}{
		{
			name: "Basic test",
			a:    []int{65, 3, 58, 678, 64},
			b:    []int{64, 2, 3, 43},
			expected: struct {
				found       bool
				intersected []int
			}{true, []int{64, 3}},
		},
		{
			name: "No intersection",
			a:    []int{1, 2, 3},
			b:    []int{4, 5, 6},
			expected: struct {
				found       bool
				intersected []int
			}{false, []int{}},
		},
		{
			name: "Empty slices",
			a:    []int{},
			b:    []int{},
			expected: struct {
				found       bool
				intersected []int
			}{false, []int{}},
		},
		{
			name: "One empty slice",
			a:    []int{1, 2, 3},
			b:    []int{},
			expected: struct {
				found       bool
				intersected []int
			}{false, []int{}},
		},
		{
			name: "Full intersection",
			a:    []int{1, 2, 3},
			b:    []int{3, 2, 1},
			expected: struct {
				found       bool
				intersected []int
			}{true, []int{3, 2, 1}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			found, intersected := FindIntersections(tt.a, tt.b)
			if found != tt.expected.found || !reflect.DeepEqual(intersected, tt.expected.intersected) {
				t.Errorf("FindIntersections(%v, %v) = (%v, %v), expected (%v, %v)",
					tt.a, tt.b, found, intersected, tt.expected.found, tt.expected.intersected)
			}
		})
	}
}
