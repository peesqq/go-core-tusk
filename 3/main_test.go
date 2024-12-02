package main

import (
	"testing"
)

func TestStringIntMap(t *testing.T) {
	m := NewStringIntMap()

	m.Add("key1", 100)
	if v, exists := m.Get("key1"); !exists || v != 100 {
		t.Errorf("Expected key1 to have value 100, but got %v (exists: %v)", v, exists)
	}

	m.Remove("key1")
	if _, exists := m.Get("key1"); exists {
		t.Errorf("Expected key1 to be removed, but it still exists")
	}

	m.Add("key2", 200)
	if !m.Exists("key2") {
		t.Errorf("Expected key2 to exist")
	}
	if m.Exists("key3") {
		t.Errorf("Expected key3 to not exist")
	}

	m.Add("key3", 300)
	copiedMap := m.Copy()
	if len(copiedMap) != 2 {
		t.Errorf("Expected copied map to have 2 elements, but got %d", len(copiedMap))
	}
	copiedMap["key4"] = 400 //независимость
	if _, exists := m.Get("key4"); exists {
		t.Errorf("Expected key4 to not exist in the original map")
	}
}
