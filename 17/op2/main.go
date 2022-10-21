// Реализовать бинарный поиск встроенными методами языка.

package main

import (
	"errors"
	"fmt"
)

func bs(slice []int, toFind int, l int, r int) (int, error) {
	for l <= r {
		piv := (l + r) / 2
		if slice[piv] == toFind {
			return piv, nil
		}

		if slice[piv] > toFind {
			r = piv - 1
		} else {
			l = piv + 1
		}
	}

	return -1, errors.New("there is no such number")
}

func Search(slice []int, toFind int) (int, error) {
	return bs(slice, toFind, 0, len(slice)-1)
}

func main() {
	slice := []int{0, 1, 11, 15, 32, 75, 201, 501}
	fmt.Println(Search(slice, -1))
	for _, num := range slice {
		fmt.Println(Search(slice, num))
	}
	fmt.Println(Search(slice, 11111))
}
