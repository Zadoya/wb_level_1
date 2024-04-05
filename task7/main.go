package main

// Реализовать конкурентную запись данных в map

import (
	"fmt"
	"sync"
)

// Применение sync.Map имеет смысл если у вас высоконагруженная 
// система с большим количеством ядер процессора (32+),
// В остальных случаях, sync.Map особо не нужен, 
// вместо него целесообразнее использовать связку map+sync.RWMutex

type MyMap struct {
	mux   sync.RWMutex
	store map[int]string
}

func NewMap() *MyMap {
	return &MyMap{mux: sync.RWMutex{},
		store: make(map[int]string),
	}
}

func (m *MyMap) Set(key int, value string) {
	m.mux.Lock()
	m.store[key] = value
	m.mux.Unlock()
}

func (m *MyMap) Get(key int) (string, bool) {
	m.mux.RLock()
	value, ok := m.store[key]
	m.mux.RUnlock()

	return value, ok
}

func main() {
	m := NewMap()
	for i := 0; i < 5; i++ {
		m.Set(i, "some string "+fmt.Sprint(i+1))
	}

	for i := range m.store {
		if result, ok := m.Get(i); ok {
			fmt.Printf("map[%d] = %s\n", i, result)
		}
	}
}

/********** 2 вариант - sync.Map *****************
func main() {
	m := sync.Map{}

	for i := 0; i < 5; i++ {
		m.Store(i, "some string " + fmt.Sprint(i + 1))
	}

	for i := 0; i < 5; i++ {
		if result, ok := m.Load(i); ok {
			fmt.Printf("map[%d] = %s\n", i, result)
		}
	}
}
*/
