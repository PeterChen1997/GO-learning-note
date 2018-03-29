package main

import "fmt"

type Base struct {
	Name string
}

func (base *Base) Foo() {}
func (base *Base) Bar() {}

type Foo struct {
	Base
}

// 从Base类继承并改写Bar方法
func (foo *Foo) Bar() {
	foo.Base.Bar()
}

func main() {
	var a Integer = 1
	if a.Less(2) {
		fmt.Println(a, "Less than 2")
	}
	a.Add(2)
	fmt.Println("a =", a)
}
