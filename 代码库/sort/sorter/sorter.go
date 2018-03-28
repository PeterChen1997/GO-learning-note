package main

// 快速解析命令行参数flag包
import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"sort/sorter/algorithms/bubblesort"
	"sort/sorter/algorithms/qsort"
	"strconv"
	"time"
)

// 下面变量为指针
var infile = flag.String("i", "infile", "File contains values for sorting")
var outfile = flag.String("o", "outfile", "File to receive sorted values")
var algorithm = flag.String("a", "qsort", "Sort algorithm")

func readValues(infile string) (values []int, err error) {
	file, err := os.Open(infile)
	if err != nil {
		fmt.Println("Failed to open the input file ", infile)
		return
	}

	defer file.Close()

	br := bufio.NewReader(file)

	values = make([]int, 0)

	for {
		// 读取
		line, isPrefix, err1 := br.ReadLine()

		if err1 != nil {
			if err1 != io.EOF {
				err = err1
			}
			break
		}

		if isPrefix {
			fmt.Println("A too long line ,seems unexpected.")
			return
		}

		str := string(line) // 字符数组 => 字符串

		// a to int
		value, err1 := strconv.Atoi(str) // 字符串 => 整型数字

		if err1 != nil {
			err = err1
			return
		}

		// 添加进values数组
		values = append(values, value)

	}
	return
}

func writeValues(values []int, outfile string) error {
	file, err := os.Create(outfile)
	if err != nil {
		fmt.Println("Failed to create the output file ", outfile)
		return err
	}

	defer file.Close()

	for _, value := range values {
		// int to a
		str := strconv.Itoa(value)
		file.WriteString(str + "\n")
	}
	return nil
}

func main() {
	flag.Parse()

	if infile != nil {
		// 需要将指针指向的变量取出来
		fmt.Println("infile = ", *infile, "outfile =", *outfile, "algorithm =", *algorithm)
	}

	values, err := readValues(*infile)
	if err == nil {
		fmt.Println("read values:", values)
		t1 := time.Now()
		switch *algorithm {
		case "qsort":
			qsort.QuickSort(values)
		case "bubblesort":
			bubblesort.BubbleSort(values)
		default:
			fmt.Println("Sorting algorithm", *algorithm, "is either unknown or unsupported.")
		}
		t2 := time.Now()

		fmt.Println("The sorting process costs", t2.Sub(t1), "to complete.")

		writeValues(values, *outfile)
	} else {
		fmt.Println(err)
	}
}
