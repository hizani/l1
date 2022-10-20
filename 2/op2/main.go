// Написать программу, которая конкурентно рассчитает значение квадратов
// чисел взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.

package main

import (
	"fmt"
	"sync"
)

func main() {
	// Создаем массив
	base := [5]int32{2, 4, 6, 8, 10}
	out := [5]int32{}

	// Инициализируем wg
	var wg sync.WaitGroup
	wg.Add(len(base))

	// Проходимся по массиву
	for idx, num := range base {
		// Передаем индекс и элемент по значению в горутину
		go func(idx int, num int32) {
			// Сохраняем результат в новый массив
			out[idx] = num * num
			// уменьшаем счетчик wg
			wg.Done()
		}(idx, num)
	}
	// Ждем завершения всех горутин
	wg.Wait()
	// Выводим массив
	fmt.Println(out)
}
