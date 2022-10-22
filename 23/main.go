// Удалить i-ый элемент из слайса.

package main

import "fmt"

func Remove[T any](slice []T, place int) []T {
	return append((slice)[:place], (slice)[place+1:]...)
}

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("before\t", slice)
	slice = Remove(slice, 4)
	fmt.Println("after\t", slice)
}
