package main

import "fmt"

type Person struct {
	Name string
}

func (Person Person) getName() {
	fmt.Println("getName()", Person.Name, "is Good Man")
}

func (Person Person) calculate() {
	result := 0
	for i := 1; i <= 1000; i++ {
		result += i
	}
	fmt.Println(Person.Name, "calculate() result", result)
}

func (Person Person) calculate2(number int) {
	result := 0
	for i := 1; i <= number; i++ {
		result += i
	}
	fmt.Println(Person.Name, "calculate2(number int) result", result)
}

func main() {
	weimin := Person{"WeiMin"}
	weimin.getName()
	weimin.calculate()
	weimin.calculate2(100)
}
