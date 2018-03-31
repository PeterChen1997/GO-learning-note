# channel 语法

## 基本语法

```go
// 声明
// chan为关键字，ElementType指定了channel所能传递的元素类型
var chanName chan ElementType

// 定义
// 定义channel使用make()
ch := make(chan int)

// 写入 - 阻塞
ch <- value
// 读取 - 阻塞
value := <-ch
```

## select

与switch非常类似，在发生IO动作时调用

最大的限制是每个case语句里必须是一个IO操作

> select 的特点是只要其中一个 case 已经完成,程序就会继续往下执行,而不会考虑其他 case 的情况

```go
select {
  case <- chan1:
  // 如果chan1成功读到数据，则进行该case处理语句
  case chan2 <- 1:
  // 如果成功向chan2写入数据，则进行该case处理语句
  default:
  // 如果上面都没有成功，则进入dafault处理
}

// 例子：死循环，随机输出0 1
ch := make(chan int, 1)
for {
  select {
  case ch <- 0:
  case ch <- 1:
  }
  i := <-ch
  fmt.Println("Value received:", i)
}

```

## 缓冲机制

```go
// 创建一个带缓冲的channel
c := make(chan int, 1024)
```

## 超时机制

利用channel的select机制，判断1s超时

```go
func main() {
	ch := make(chan int, 1)
	timeout := make(chan bool, 1)

	// 是否进行IO
	// ch <- 1

	go func() {
		time.Sleep(1e9) // 等待1秒钟，默认为ns为单位
		timeout <- true
	}()
	
	// 然后我们把timeout这个channel利用起来
	select {
	case <-ch:
		// 从ch中读取到数据
		fmt.Println("从ch中读取到数据")
	case <-timeout:
		// 一直没有从ch中读取到数据,但从timeout中读取到了数据
		fmt.Println("一直没有从ch中读取到数据,但从timeout中读取到了数据")
	}
}
```

## channel的传递

使用channel也能通过channel传递的特性实现pipe（管道）

```go
type PipeData struct {
  value int
  handler func(int) int
  next chan int
}

func handle(queue chan *PipeData) {
  for data := range queue {
    data.next <- data.handler(data.value)
  }
}
```

## 单向channel

单向channel只能发送或者接收数据，其实是对channel的一种使用限制

```go
// 正常channel
var ch1 chan int 
// 单向channel，只用于写float64数据
var ch2 chan<- float64
// 单向channel，只用于读取int数据
var ch3 <-chan int

// 类型转换
ch4 := make(chan int)
ch5 := <-chan int(ch4) // 单向读取
ch6 := chan<- int(ch4) // 单项写入

// 例子
func Parse(ch <-chan int) {
  for value := range ch {
    fmt.Println("Parsing value", value)
  }
}
```

## 关闭channel

```go
close(ch)

// 判断是否关闭
x, ok := <-ch
// 看第二个bool返回值即可，如果为false则表示ch 已经被关闭
```