package main

import "fmt"

type Integer int

type Rect struct {
	x, y float64,
	width, height float64
}

// 面向对象的用法
func (a Integer) Less(b Integer) bool {
	return a < b
}
func (a *Integer) Add(b Integer) {
	*a += b
}
// 自定义结构体类型的构造函数
func NewRect(x, y, width, height float64) {
	return &Rect{x, y, width, height}
}

func main() {
	var a Integer = 1
	if a.Less(2) {
		fmt.Println(a, "Less than 2")
	}
	a.Add(2)
	fmt.Println("a =", a)
}
