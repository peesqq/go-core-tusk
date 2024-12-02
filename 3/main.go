package main

import "sync"

type StringIntMap struct {
	data map[string]int
	mu   sync.RWMutex
}

func NewStringIntMap() *StringIntMap {
	return &StringIntMap{
		data: make(map[string]int),
	}
}

func (m *StringIntMap) Add(key string, value int) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.data[key] = value
}

func (m *StringIntMap) Remove(key string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	delete(m.data, key)
}

func (m *StringIntMap) Copy() map[string]int {
	m.mu.RLock()
	defer m.mu.RUnlock()

	newMap := make(map[string]int)
	for k, v := range m.data {
		newMap[k] = v
	}
	return newMap
}

func (m *StringIntMap) Exists(key string) bool {
	m.mu.RLock()
	defer m.mu.RUnlock()
	_, exists := m.data[key]
	return exists
}

func (m *StringIntMap) Get(key string) (int, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	value, exists := m.data[key]
	return value, exists
}
