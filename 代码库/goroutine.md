# goroutine

## 例子

```go
package main

import "fmt"

func Add(x, y int) {
	z := x + y
	// 无法看到输出
	// golang 在 main() 函数结束时返回，并不等待其他goroutine
	fmt.Println(z)
}

func main() {
	for i := 0; i < 10; i++ {
		go Add(i, i)
	}
}
```

## 解决方法

```go
package main

import "fmt"

func Count(ch chan int) {
	fmt.Println("Counting")
	// 函数调用后，向channel中写入一个值，写入阻塞
	ch <- 1
}

func main() {
	chs := make([]chan int, 10)
	for i := 0; i < 10; i++ {
		chs[i] = make(chan int)
		// 分配channel给10个不同的goroutine
		go Count(chs[i])
	}

	for _, ch := range chs {
		// 在channel被读取前，是阻塞的
		<-ch
	}
}

```