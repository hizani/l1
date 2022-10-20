// Реализовать конкурентную запись данных в map.

package main

import (
	"fmt"
	"math/rand"
	"sync"
)

type ThreadsafeMap struct {
	Map   map[int]int
	Mutex sync.Mutex
}

func (m *ThreadsafeMap) Set(key int, value int) {
	m.Mutex.Lock()
	m.Map[key] = value
	m.Mutex.Unlock()
}

func main() {
	m := ThreadsafeMap{Map: make(map[int]int)}
	wg := sync.WaitGroup{}
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			for i := 0; i < 100; i++ {
				m.Set(rand.Intn(100), rand.Intn(100))
			}
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("map:", m.Map)
}
