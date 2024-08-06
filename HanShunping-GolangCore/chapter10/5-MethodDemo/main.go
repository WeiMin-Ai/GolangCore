package main

import "fmt"

type Person struct {
	Name string
}

func (p Person) changeName(name string) {
	p.Name = name
	fmt.Println("changeName() Name:", p.Name)
}

func (p Person) calculate() {
	res := 0
	for i := 1; i <= 1000; i++ {
		res += i
	}
	fmt.Println("calculate() res:", res)
}

func (p Person) calculate2(number int) {
	res := 0
	for i := 1; i <= number; i++ {
		res += i
	}
	fmt.Println("calculate2() res:", res)
}

func (p Person) getSum(number1, number2 int) int {
	return number1 + number2
}

func main() {
	var person Person
	person.Name = "WeiMin"
	person.changeName("WeiMin123")
	person.calculate()
	person.calculate2(1000)
	fmt.Println(person.getSum(1000, 1000))
}
