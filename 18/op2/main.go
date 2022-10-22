// Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде.
// По завершению программа должна выводить итоговое значение счетчика.

package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// Используем atomic для предотвращения половинчатого инкремента
// в конкурентной среде
type Counter struct {
	counter int32
}

func (c *Counter) Increment() {
	atomic.AddInt32(&c.counter, 1)
}

func (c *Counter) Get() int32 {
	return atomic.LoadInt32(&c.counter)
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
