// Дана последовательность температурных колебаний: -25.4, -27.0 13.0, 19.0, 15.5, 24.5, -21.0, 32.5.
// Объединить данные значения в группы с шагом в 10 градусов. Последовательность в подмножноствах не важна.
// Пример: -20:{-25.0, -27.0, -21.0}, 10:{13.0, 19.0, 15.5}, 20: {24.5}, etc.

package main

import "fmt"

type TempSlice []float64

type TempMap map[float64]TempSlice

func NewTempSet() TempMap {
	return make(TempMap)
}

// То же решение, если нам все рано на уникальность элементов
func (tm TempMap) Add(temp float64) {
	tenthTemp := float64(int(temp/10) * 10.0)
	tm[tenthTemp] = append(tm[tenthTemp], temp)
}

func main() {
	base := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	tempGroup := NewTempSet()

	for _, temp := range base {
		tempGroup.Add(temp)
	}
	fmt.Printf("%v\n", tempGroup)
}
