// Разработать программу, которая проверяет, что все символы в строке уникальные (true — если уникальные, false etc).
// Функция проверки должна быть регистронезависимой.
// Например:
// abcd — true
// abCdefAaf — false
// aabcd — false

package main

import (
	"fmt"
	"strings"
)

func IsUnique1(text string) bool {
	text = strings.ToLower(text)

	for _, r := range text {
		if strings.Count(text, string(r)) > 1 {
			return false
		}
	}
	return true
}

func IsUnique2(text string) bool {
	text = strings.ToLower(text)

	m := make(map[rune]bool)
	for _, v := range text {
		if m[v] {
			return false
		}
		m[v] = true
	}
	return true
}

func main() {
	str1 := "abcd"
	str2 := "abCdefAaf"
	str3 := "aabcd"

	fmt.Println(str1, IsUnique1(str1))
	fmt.Println(str2, IsUnique1(str2))
	fmt.Println(str3, IsUnique1(str3))
	fmt.Println("=======")
	fmt.Println(str1, IsUnique2(str1))
	fmt.Println(str2, IsUnique2(str2))
	fmt.Println(str3, IsUnique2(str3))
}
