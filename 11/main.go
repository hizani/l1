// Реализовать пересечение двух неупорядоченных множеств.

package main

import "fmt"

type Set map[int]struct{}

func NewSet() Set {
	return make(Set)
}

func GetIntersection(s1 Set, s2 Set) Set {
	interSet := make(Set)

	// Находим меньшее из множеств чтобы меньше итерироваться
	smallest := s1
	biggest := s2
	if len(s2) < len(s1) {
		interSet = s2
		biggest = s1
	}

	// Находим пересечения
	for key := range smallest {
		if _, found := biggest[key]; found {
			interSet[key] = struct{}{}

		}
	}
	return interSet
}

func main() {
	nums1 := []int{1, 2, 4, 5}
	nums2 := []int{-3, -2, -1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	set1 := NewSet()
	set2 := NewSet()
	for _, v := range nums1 {
		set1[v] = struct{}{}
	}
	for _, v := range nums2 {
		set2[v] = struct{}{}
	}
	newSet := GetIntersection(set1, set2)
	fmt.Println(newSet)
}
