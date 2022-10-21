// Дана последовательность температурных колебаний: -25.4, -27.0 13.0, 19.0, 15.5, 24.5, -21.0, 32.5.
// Объединить данные значения в группы с шагом в 10 градусов. Последовательность в подмножноствах не важна.
// Пример: -20:{-25.0, -27.0, -21.0}, 10:{13.0, 19.0, 15.5}, 20: {24.5}, etc.

package main

import (
	"fmt"
	"math"
)

// Подобие множества температур
type TempSet map[float64]struct{}

// ключ - шаг в 10 градусов, значение - множество температур
type TempMap map[float64]TempSet

func NewTempSet() TempMap {
	return make(TempMap)
}

func (tm TempMap) Add(temp float64) {
	// получаем шаг
	tenthTemp := math.Trunc(temp/10) * 10.0
	// Если множество еще не создано, то создаем
	if tm[tenthTemp] == nil {
		tm[tenthTemp] = make(TempSet)
	}
	// Добавляем число
	tm[tenthTemp][temp] = struct{}{}
}

func main() {
	base := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	tempGroup := NewTempSet()

	for _, temp := range base {
		tempGroup.Add(temp)
	}
	fmt.Printf("%v\n", tempGroup)
}
