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

	// Инициализируем wg
	wg := sync.WaitGroup{}
	wg.Add(len(base))

	// Проходимся по массиву
	for i := range base {
		// Передаем указатель на
		// элемент массива в горутину
		go func(num *int32) {
			// Изменяем значение элемента массива по указателю
			*num = (*num) * (*num)
			// уменьшаем счетчик wg
			wg.Done()
		}(&base[i])
	}
	// Ждем завершения всех горутин
	wg.Wait()
	// Выводим измененный массив
	fmt.Println(base)
}