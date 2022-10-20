// Дана последовательность чисел: 2,4,6,8,10.
// Найти сумму их квадратов(22+32+42….) с использованием конкурентных вычислений.

package main

import (
	"fmt"
	"sync"
)

func main() {
	// Создаем массив
	base := [5]int32{2, 4, 6, 8, 10}
	var acc int32

	wg := sync.WaitGroup{}
	wg.Add(5)

	// Проходимся по массиву
	for _, num := range base {
		// Передаем элемент по значению и ссылку на аккомулятор в горутину
		go func(num int32, acc *int32) {
			// Сохраняем результат в новый массив
			*acc += num * num
			// уменьшаем счетчик wg
			wg.Done()
		}(num, &acc)
	}
	// Ждем завершения всех горутин
	wg.Wait()
	// Выводим массив
	fmt.Println(acc)
}
