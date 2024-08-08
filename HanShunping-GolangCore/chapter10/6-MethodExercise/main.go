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

func (Person Person) getSum(number1, number2 int) int {
	return number1 + number2
}

type Circle struct {
	radius float64
}

func (Circle *Circle) area() float64 {
	return 3.14 * Circle.radius * Circle.radius
}

func main() {
	weimin := Person{"WeiMin"}
	weimin.getName()
	weimin.calculate()
	weimin.calculate2(20)
	res := weimin.getSum(1, 2)
	fmt.Println("Person.getSum() Result:", res)

	circle := Circle{radius: 11.0}
	res2 := circle.area()
	fmt.Println("Circle.getSum() Result:", res2)
}
