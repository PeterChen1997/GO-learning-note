package main

import (
	"fmt"
	"strconv"
)

type Element interface{}
type List []Element

type Person struct {
	name string
	age  int
}

func (p Person) String() string {
	return "(name: " + p.name + " - age: " + strconv.Itoa(p.age) + " years)"
}

func main() {
	list := make(List, 3)
	list[0] = 1       // int
	list[1] = "Hello" // string
	list[2] = Person{"Dennis", 70}

	for index, element := range list {
		switch value := element.(type) {
		case int:
			fmt.Printf("int")
		case string:
			fmt.Printf("string")
		case Person:
			fmt.Printf("Person")
		default:
			fmt.Printf("Unknow")
		}
	}
}
