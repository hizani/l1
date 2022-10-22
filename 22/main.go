// Разработать программу, которая перемножает, делит, складывает,
// вычитает две числовых переменных a,b, значение которых > 2^20.

package main

import (
	"fmt"
	"math/big"
)

func main() {
	var a int64 = 9_146_744_073_709_551_615
	var b int64 = 9_146_744_073_709_551_616
	fmt.Printf("a: %d\nb: %d\n", a, b)

	var bigInt big.Int
	fmt.Println("a + b = ", bigInt.Add(big.NewInt(a), big.NewInt(b)))

	fmt.Println("a - b = ", bigInt.Sub(big.NewInt(a), big.NewInt(b)))

	fmt.Println("a / b = ", bigInt.Div(big.NewInt(a), big.NewInt(b)))

	fmt.Println("a * b = ", bigInt.Mul(big.NewInt(a), big.NewInt(b)))

}
