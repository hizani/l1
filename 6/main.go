// Реализовать все возможные способы остановки выполнения горутины.

package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func WaitGroupGoroutine(wg *sync.WaitGroup) {
	wg.Wait()
	fmt.Println("WaitGroupGoroutine finished")

}

func ContextGoroutine(ctx context.Context) {
	<-ctx.Done()
	fmt.Println("ContextGoroutine finished")
}

func ChanGoroutine(c chan interface{}) {
	<-c
	fmt.Println("ChanGoroutine finished")
}

func ForChanGoroutine(c chan int) {
	for range c {
		// что то делает
	}
	fmt.Println("ForChanGoroutine finished")
}

func main() {
	// 1. Завершаем горутину уменьшая счетчик wg
	wg := sync.WaitGroup{}
	wg.Add(1)
	go WaitGroupGoroutine(&wg)
	wg.Done()

	// 2. Завершаем горутину вызывая cancel() у ctx
	ctx, cancel := context.WithCancel(context.Background())
	go ContextGoroutine(ctx)
	cancel()

	// 3. Завершаем горутину читая из канала
	c1 := make(chan interface{})
	go ChanGoroutine(c1)
	c1 <- "kill"

	// 4. Завершаем горутину закрывая канал
	c2 := make(chan int)
	go ForChanGoroutine(c2)
	for i := 0; i < 5; i++ {
		c2 <- i
	}
	close(c2)

	time.Sleep(time.Millisecond * 500)

}
