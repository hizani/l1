// Разработать программу, которая переворачивает слова в строке.
// Пример: «snow dog sun — sun dog snow».

package main

import (
	"fmt"
	"strings"
)

func main() {
	base := "snow dog sun"

	temp := strings.Fields(base)
	builder := strings.Builder{}
	for i := len(temp) - 1; i >= 0; i-- {
		builder.WriteString(temp[i] + " ")
	}
	new := strings.TrimSpace(builder.String())
	fmt.Println(new)
}
