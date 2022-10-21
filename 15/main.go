// К каким негативным последствиям может привести данный фрагмент кода, и как это исправить?
// Приведите корректный пример реализации.

package main

import (
	"fmt"
	"math/rand"
	"time"
)

var justString string

func createHugeString(size int) (str string) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		str += string(rune(rand.Intn(145000)))
	}
	return
}

// плохая функция
func someFunc() {
	v := createHugeString(1 << 10)
	// Получаем не 100 символов, а 100 байт строки.
	// Это может привести к неполному сохранению
	// символов больше одного байта в конце слайса
	justString = v[:100]
}

// хорошая функция
func someGoodFunc() {
	v := createHugeString(1 << 10)
	// конвертируем строку в []rune, чтобы нарезать на символы
	r := []rune(v)
	justString = string(r[:100])
}

func main() {
	someFunc()
	fmt.Println(justString)
	someGoodFunc()
	fmt.Println(justString)
}
