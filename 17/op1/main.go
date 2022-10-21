// Реализовать бинарный поиск встроенными методами языка.

package main

import (
	"errors"
	"fmt"
)

// Через нарезку слайса
func BinarySearch(slice []int, toFind int) (int, error) {
	l := 0
	r := len(slice)

	if l >= r {
		return -1, errors.New("there is no such number")
	}

	piv := (r + l) / 2

	if slice[piv] == toFind {
		return piv, nil
	}
	if slice[piv] > toFind {
		return BinarySearch(slice[:piv], toFind)
	}

	num, err := BinarySearch(slice[piv+1:], toFind)
	return piv + 1 + num, err
}

func main() {
	slice := []int{0, 1, 11, 15, 32, 75, 201, 501}
	fmt.Println(BinarySearch(slice, -1))
	for _, num := range slice {
		fmt.Println(BinarySearch(slice, num))
	}
	fmt.Println(BinarySearch(slice, 11111))
	// BinarySearch(slice, 11)
}
