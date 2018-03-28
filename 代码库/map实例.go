package main

import "fmt"

type PersonInfo struct {
	ID      string
	Name    string
	Address string
}

func main() {
	// personDB 是声明的变量名，并未赋值
	var personDB map[string]PersonInfo
	// 通过make()来创建一个新的map,并赋值给personDB
	// make()第三个参数为初始存储能力
	personDB = make(map[string]PersonInfo)
	personDB["123"] = PersonInfo{"123", "Tom", "Room213"}
	// 上面的赋值等同于
	// personDB = map[string]PersonInfo{
	// 	"123": PersonInfo{"123", "Tom", "Room213"},
	// }

	person, OK := personDB["123"]

	if OK {
		fmt.Println("Found person", person.Name, "with ID", person.ID)
	} else {
		fmt.Println("Didn't found someone.")
	}
}
