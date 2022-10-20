// Дана структура Human (с произвольным набором полей и методов).
// Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования).

package main

import "fmt"

// Родительская структура
type Human struct {
	name string
	age  int32
}

func (Human) human_echo() {
	fmt.Println("I'm human!")
}

func (h Human) human_info() {
	fmt.Printf("Name: %s Age: %d\n", h.name, h.age)
}

// 1
type Action1 struct {
	// Анонимная копия "родителя"
	Human
}

// 2
type Action2 struct {
	// Анонимный указатель на "родителя"
	*Human
}

func main() {
	human := Human{"Абоба", 10}
	action_copy := Action1{human}
	action_pointer := Action2{&human}

	fmt.Println("Копия родителя")
	action_copy.human_echo()
	action_copy.human_info()
	fmt.Println("-------------------------")
	fmt.Println("Указатель на родителя")
	action_pointer.human_echo()
	action_pointer.human_info()
}
