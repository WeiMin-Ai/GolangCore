// 练习：编写方法，使给定的一个二维数组(3x3)转置

package main

import "fmt"

type MethodUtils struct {
	Arrays [3][3]int
}

func (MethodUtils *MethodUtils) ArrayTranspose() {
	var tempArrays [3][3]int

	length := len(MethodUtils.Arrays[0])
	for i := 0; i < length; i++ {
		for j := 0; j < length; j++ {
			tempArrays[i][j] = MethodUtils.Arrays[j][i]
		}
	}
	MethodUtils.Arrays = tempArrays
}

func main() {
	var methodUtils MethodUtils
	methodUtils.Arrays[0] = [3]int{1, 2, 3}
	methodUtils.Arrays[1] = [3]int{4, 5, 6}
	methodUtils.Arrays[2] = [3]int{7, 8, 9}

	methodUtils.ArrayTranspose()

	for _, array := range methodUtils.Arrays {
		fmt.Println(array)
	}

}
