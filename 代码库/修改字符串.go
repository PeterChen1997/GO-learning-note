package main

import (
	"fmt"
)

func main() {
	s := "Hello World"
	c := []byte(s) // 转换为[]byte
	c[0] = 's'
	s2 := string(c) // 转换为string
	fmt.Println(s2)
}
