// Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0.

package main

import (
	"fmt"
)

func ChangeBit(num *int64, pos uint) {
	// XOR устанавливает бит на 1 только
	// если один из битов равен 1
	*num ^= (1 << pos)
}

func main() {
	var num int64 = 31
	fmt.Printf("%.64b\n", num)
	ChangeBit(&num, 2)
	fmt.Printf("%.64b\n", num)
}
