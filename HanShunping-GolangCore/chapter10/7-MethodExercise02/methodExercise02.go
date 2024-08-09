// 方法练习

package main

import "fmt"

type MethodUtils struct {
}

type Calculator struct {
	Sign             string
	Number1, Number2 float64
	Result           float64
}

func (MethodUtils *MethodUtils) PrintRectangle() {
	for i := 0; i < 8; i++ {
		for j := 0; j < 10; j++ {
			fmt.Print("*  ")
		}
		fmt.Println()
	}
}

func (MethodUtils *MethodUtils) PrintRectangle2(m, n int) {
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			fmt.Print("*  ")
		}
		fmt.Println()
	}
}

func (MethodUtils *MethodUtils) RectangleCalculation(length, width float64) float64 {
	return length * width
}

func (MethodUtils *MethodUtils) IsEven(number int) bool {
	if number%2 == 0 {
		return true
	}
	return false
}

func (MethodUtils *MethodUtils) Print(n, m int, char string) {
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			fmt.Printf("%s ", char)
		}
		fmt.Println()
	}
}

func (Calculator *Calculator) Calculate() {
	//if Calculator.Sign == "" || Calculator.Number1 == 0 || Calculator.Number2 == 0 {
	//	return
	//}

	switch Calculator.Sign {
	case "+":
		Calculator.Result = Calculator.Number1 + Calculator.Number2
	case "-":
		Calculator.Result = Calculator.Number1 - Calculator.Number2
	case "/":
		Calculator.Result = Calculator.Number1 / Calculator.Number2
	case "*":
		Calculator.Result = Calculator.Number1 * Calculator.Number2
	}
}

func main() {
	var methodUtils MethodUtils
	methodUtils.PrintRectangle()
	methodUtils.PrintRectangle2(5, 6)
	res := methodUtils.RectangleCalculation(5.4, 6)
	fmt.Println("Result:", res)
	fmt.Println("IsOddAndEven:", methodUtils.IsEven(2))
	methodUtils.Print(5, 5, "\\*\\")

	var cal Calculator = Calculator{Sign: "+", Number1: 5, Number2: 10, Result: 0}
	cal.Calculate()
	fmt.Printf("Result: %v", fmt.Sprintf("%.2f", cal.Result))
}
