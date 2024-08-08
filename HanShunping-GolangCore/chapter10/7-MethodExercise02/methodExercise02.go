// 方法练习题目

package main

import "fmt"

type MethodUtils struct {
}

func (MethodUtils MethodUtils) printRectangle() {
	for i := 0; i < 8; i++ {
		for j := 0; j < 10; j++ {
			fmt.Print("*  ")
		}
		fmt.Println()
	}
}

func (MethodUtils MethodUtils) printRectangle2(m, n int) {
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			fmt.Print("*  ")
		}
		fmt.Println()
	}
}

func main() {
	var methodUtils MethodUtils
	//methodUtils.printRectangle()
	methodUtils.printRectangle2(5, 6)
}
