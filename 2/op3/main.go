// Написать программу, которая конкурентно рассчитает значение квадратов
// чисел взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.

package main

import (
	"fmt"
)

// используем шаблон worker
func main() {
	worker := func(jobs <-chan int32, result chan<- int32) {
		// берем числа из jobs
		for n := range jobs {
			// передаем результаты вычисления в result
			result <- n * n
		}
	}

	// Создаем массив
	base := [5]int32{2, 4, 6, 8, 10}
	// Создаем два канала: один принимает результат вычисления,
	// другой отдает числа для вычисления
	jobs := make(chan int32, len(base))
	result := make(chan int32, len(base))

	// Создаем 5 работников
	for range base {
		go worker(jobs, result)
	}

	// Передаем работникам элементы base
	for _, num := range base {
		jobs <- num
	}
	close(jobs)

	for range base {
		fmt.Printf("%d ", <-result)
	}
	fmt.Println()
	close(result)
}
