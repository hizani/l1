// Реализовать быструю сортировку массива (quicksort) встроенными методами языка.

package main

import "fmt"

func Sort(slice []int) {
	quickSort(0, len(slice)-1, slice)
}

func quickSort(l int, r int, slice []int) {
	minPtr := l
	maxPtr := r
	piv := slice[(maxPtr+minPtr)/2]

	for minPtr <= maxPtr {
		for slice[minPtr] < piv {
			minPtr++
		}
		for slice[maxPtr] > piv {
			maxPtr--
		}
		if minPtr <= maxPtr {
			temp := slice[minPtr]
			slice[minPtr] = slice[maxPtr]
			slice[maxPtr] = temp
			minPtr++
			maxPtr--
		}
	}

	if l < maxPtr {
		quickSort(l, maxPtr, slice)
	}
	if r > minPtr {
		quickSort(minPtr, r, slice)
	}
}

func main() {
	array := [...]int{0, 5, 8, 3, 5, 9, 7, 3, 6, 8, 6, 7, 2, 4, 1, 5, 8, 7, 0, 8, 0, 43, 6, 3}
	fmt.Println(array)
	Sort(array[:])
	fmt.Println(array)
}
