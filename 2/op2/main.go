// Написать программу, которая конкурентно рассчитает значение квадратов
// чисел взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.

package main

import (
	"fmt"
	"sync"
)

func main() {
	base := [5]int32{2, 4, 6, 8, 10}
	out := [5]int32{}

	var wg sync.WaitGroup
	wg.Add(len(base))

	for idx, num := range base {
		go func(idx int, num int32) {
			out[idx] = num * num
			wg.Done()
		}(idx, num)
	}
	wg.Wait()
	fmt.Println(out)
}
