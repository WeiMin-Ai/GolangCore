package main

import (
	"GolangCore/calcproj/src/simplemath"
	"fmt"
	"os"
	"strconv"
)

var Usage = func() {
	fmt.Println("Usage: calc command [arguments] ...")
	fmt.Println("The commands are:")
	fmt.Println("sqrt\tSquart root of a non-negative value.")
	fmt.Println("add\tAddition of two values.")
}

func main() {
	args := os.Args
	if args == nil || len(args) < 3 {
		Usage()
		return
	}
	switch args[1] {
	case "add":
		if len(args) != 4 {
			fmt.Println("Usage: calc add <number1> <number1>")
			return
		}
		v1, err1 := strconv.Atoi(args[2])
		v2, err2 := strconv.Atoi(args[3])
		if err1 != nil || err2 != nil {
			fmt.Println("Usage: calc add <number1> <number1>")
			return
		}
		ret := simplemath.Add(v1, v2)
		fmt.Println("Result:", ret)
	case "sqrt":
		if len(args) != 3 {
			fmt.Println("Usage: calc sqrt <number>")
			return
		}
		v1, err1 := strconv.Atoi(args[2])
		if err1 != nil {
			fmt.Println("Usage: calc sqrt <number>")
			return
		}
		ret := simplemath.Sqrt(v1)
		fmt.Println("Result:", ret)
	default:
		Usage()
	}
}
