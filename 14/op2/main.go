// Разработать программу, которая в рантайме способна
// определить тип переменной: int, string, bool, channel из переменной типа interface{}.

package main

import "fmt"

func CheckType(i interface{}) {
	switch i.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case bool:
		fmt.Println("bool")
	case chan int:
		fmt.Println("chan")
	}
}

func main() {
	var i int
	var str string
	var b bool
	var ch1 chan int
	var ch2 chan string

	CheckType(i)
	CheckType(str)
	CheckType(b)
	CheckType(ch1)
	CheckType(ch2)

}
