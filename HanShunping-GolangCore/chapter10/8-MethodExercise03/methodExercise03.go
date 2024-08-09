// 方法练习题2

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type MethodUtils struct {
	Number int
	Exit   bool
}

// readString 获取用户输入数据
func (MethodUtils *MethodUtils) getNumber() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("输入exit则退出程序")
	fmt.Println("请输入0-9数字: ")

	input, _ := reader.ReadString('\n')
	input = strings.Split(input, "\n")[0]

	if input == "exit" {
		MethodUtils.Exit = true
	}

	number, _ := strconv.Atoi(input)

	MethodUtils.Number = number
}

// getNumber1 从键盘接受整数（1～9），并打印对应乘法表。
func (MethodUtils *MethodUtils) print() {
	// 1. 接收用户输入
	MethodUtils.getNumber()
	number := MethodUtils.Number

	// 2. 打印
	if number < 0 || number > 9 {
		fmt.Println("无效数字！请输入1~9.")
	} else if MethodUtils.Exit {
		os.Exit(0)
	} else {
		for i := 1; i <= number; i++ {
			for j := 1; j <= i; j++ {
				fmt.Printf("%v*%v=%v ", j, i, j*i)
			}
			fmt.Println("")
		}
	}
}

func main() {
	var methodUtils MethodUtils

	for {
		methodUtils.print()
	}
}
