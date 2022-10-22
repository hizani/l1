// Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде.
// По завершению программа должна выводить итоговое значение счетчика.

package main

import (
	"fmt"
	"sync"
)

// Используем Mutex для синхронизации
type Counter struct {
	counter int
	mutex   sync.Mutex
}

func (c *Counter) Increment() {
	c.mutex.Lock()
	c.counter++
	c.mutex.Unlock()
}

func (c *Counter) Get() int {
	c.mutex.Lock()
	result := c.counter
	c.mutex.Unlock()
	return result
}

func main() {
	var counter Counter
	wg := sync.WaitGroup{}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			for i := 0; i < 100; i++ {
				counter.Increment()
			}
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(counter.Get())
}
