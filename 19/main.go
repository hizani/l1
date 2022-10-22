// Разработать программу, которая переворачивает подаваемую
// на ход строку (например: «главрыба — абырвалг»). Символы могут быть unicode.

package main

import "fmt"

func main() {
	base := "главрыба"
	runed := []rune(base)
	new := ""
	for i := len(runed) - 1; i >= 0; i-- {
		new += string(runed[i])
	}
	fmt.Println(new)
}
