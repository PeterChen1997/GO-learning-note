package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var factsArr = [...]string{"有毛发", "有奶", "有羽毛", "会飞", "会下蛋", "吃肉", "有犬齿", "有爪", "眼盯前方", "有蹄", "嚼反刍", "黄褐色", "身上有暗斑点", "身上有黑色条纹", "有长脖子", "有长腿", "不会飞", "会游泳", "有黑白二色", "善飞", "哺乳动物", "鸟", "食肉动物", "蹄类动物"}

var rulesArr = [7][10]int{
	{0, 1, 13, 11, 21, 5, 6, 7, 8, 20},
	{0, 1, 12, 11, 21, 5, 6, 7, 8, 20},
	{23, 14, 15, 12, 0, 1, 9, 10, 20, -1},
	{23, 13, 9, 0, 1, 10, 20, -1, -1, -1},
	{21, 16, 14, 15, 18, 3, 2, 4, -1, -1},
	{21, 3, 2, 4, 17, 16, 18, -1, -1, -1},
	{21, 3, 2, 4, 19, -1, -1, -1, -1, -1},
}

var resultArr = [7]string{"金钱豹", "虎", "长颈鹿", "斑马", "鸵鸟", "企鹅", "信天翁"}

var inputArr []int

var err error

func checkEnd(userInput []string) bool {
	for _, item := range userInput {
		if item != "END" {
			item, _ := strconv.Atoi(item)
			inputArr = append(inputArr, item)
		} else {
			return true
		}
	}
	return false
}

func handleProcess() {
	maxIndex := -1
	maxValue := 1
	tempValue := 1
	tempArr := []int{}

	for i := 0; i < 7; i++ {
		for j := 0; j < 10; j++ {
			for _, item := range inputArr {
				if item == rulesArr[i][j] {
					tempValue++
				}
			}
		}

		if tempValue > maxValue {
			maxIndex = i
			maxValue = tempValue
			tempArr = []int{i}
		} else if tempValue == maxValue && tempValue == len(inputArr)+1 {
			// TODO:相等继续添加
			tempArr = append(tempArr, i)
		}

		tempValue = 1
	}

	// 更多处理
	if len(tempArr) > 1 {
		handleMore(tempArr)
	} else if len(tempArr) == 0 {
		fmt.Println(">>> 找不到满足要求的实物")
	} else {
		fmt.Println(">>> 最终结果为")
		fmt.Println(resultArr[maxIndex])
	}

}

func handleMore(tempArr []int) {
	fmt.Println(">>> 当前可能的结果有：")
	for _, index := range tempArr {
		fmt.Printf(resultArr[index] + " ")
	}
	fmt.Println()
	// 再次接收输入
	handleInput()
	// 在此处理
	handleProcess()
}

func handleInput() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println()
	for {
		fmt.Print(">>> 请输入包含的特征序号(输入：END结束输入) > ")
		readLine, _, _ := reader.ReadLine()
		line := string(readLine)
		tokens := strings.Split(line, " ")
		if isEnd := checkEnd(tokens); isEnd {
			fmt.Println()
			break
		}
	}
}

func main() {
	// 输出类型
	fmt.Println(">>> 下面是一些动物的特征：")
	fmt.Println()
	for index, item := range factsArr {
		fmt.Printf("|%-4v|%v\n", index, item)
	}
	// 接收输入
	handleInput()
	// 处理输入结果，输出答案
	handleProcess()
}
