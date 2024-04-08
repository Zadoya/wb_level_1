package main

// Реализовать структуру-счетчик, которая будет инкрементироваться 
// в конкурентной среде. По завершению программа должна выводить 
// итоговое значение счетчика.

import (
	"fmt"
	"sync"
)

// Cтруктура счетчика
type Counter struct {
    mutex sync.RWMutex
	count int64
}

// Функция для увеличения значения счетчика на 1
func (c *Counter) Increment(wg *sync.WaitGroup) {
	defer wg.Done()
    c.mutex.Lock()
	c.count++
	c.mutex.Unlock()
}

// Функция для получения текущего значения счетчика
func (c *Counter) GetCount() int64 {
	var result int64
		c.mutex.RLock()
		result = c.count
		c.mutex.RUnlock()
	
	return result
}
func main() {
	wg 		:= sync.WaitGroup{}
	counter := Counter{}

    for i := 0; i < 100; i++ {
		wg.Add(1)
    	go counter.Increment(&wg)
	}
	wg.Wait()
	count := counter.GetCount()
	fmt.Println("Итоговое значение счётчика: ", count) 
}

/************* реализация на атомиках ****************
import (
	"fmt"
	"sync"
	"sync/atomic"
)

// Cтруктура счетчика
type Counter struct {
	count int64
}

// Функция для увеличения значения счетчика на 1
func (c *Counter) Increment(wg *sync.WaitGroup) {
	defer wg.Done()
	atomic.AddInt64(&c.count, 1)

}

// Функция для получения текущего значения счетчика
func (c *Counter) GetCount() int64 {
	return atomic.LoadInt64(&c.count)
}
func main() {
	wg 		:= sync.WaitGroup{}
	counter := Counter{}

    for i := 0; i < 100; i++ {
		wg.Add(1)
    	go counter.Increment(&wg)
	}
	wg.Wait()
	count := counter.GetCount()
	fmt.Println("Итоговое значение счётчика: ", count) 
}
*/