// Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество.

package main

import "fmt"

type Set map[string]struct{}

func main() {
	str := []string{"cat", "cat", "dog", "cat", "tree"}
	set := make(Set)
	for _, v := range str {
		set[v] = struct{}{}
	}

	fmt.Println(set)
}
