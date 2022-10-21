// Поменять местами два числа без создания временной переменной.

package main

import "fmt"

func main() {
	x, y := 15, 20

	x = x + y // x = 35, y = 20
	y = x - y // x = 35, y = 15
	x = x - y // x = 20, y = 15

	fmt.Printf("x = %d, y = %d\n", x, y)
}
