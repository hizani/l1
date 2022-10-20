// Разработать программу, которая будет последовательно отправлять значения в канал,
// а с другой стороны канала — читать. По истечению N секунд программа должна завершаться.

package main

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {

	// берем время завершения из аргумента
	if len(os.Args) < 2 {
		fmt.Println("no arguments")
		return
	}
	sec, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	// устанавливаем контекст с таймаутом
	ctx, _ := context.WithTimeout(context.Background(), time.Second*time.Duration(sec))
	// создаем канал куда будем писать и откуда читать
	ch := make(chan int64)

	// запускаем запись в канал
	go func(ctx context.Context, ch chan int64) {
		var i int64
		for {
			select {
			case <-ctx.Done():
				close(ch)
				return
			default:
				ch <- i
			}
			i++
		}
	}(ctx, ch)

	// читаем пока есть что читать, и пока канал не закрыт
	for val := range ch {
		fmt.Println(val)
	}

}
