package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Student struct {
	Id     int
	Name   string
	Gender string
	Age    int
	Score  float64
}

type Visitor struct {
	Name string
	Age  int
	Exit bool
}

func (Visitor *Visitor) InputData() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("请输入姓名: ")
	name, _ := reader.ReadString('\n')
	name = strings.Split(name, "\n")[0]

	if name == "exit" {
		fmt.Println("退出程序。")
		os.Exit(0)
	}

	fmt.Println("请输入年龄: ")
	input, _ := reader.ReadString('\n')
	age, _ := strconv.Atoi(strings.Split(input, "\n")[0])

	Visitor.Name, Visitor.Age = name, age
}

func (Visitor *Visitor) Charge() int {
	if Visitor.Age > 18 {
		return 20
	} else {
		return 0
	}
}

func (Visitor *Visitor) Print(price int) {
	if Visitor.Age >= 90 || Visitor.Age < 8 {
		fmt.Println("年龄限制，无法进入。")
		return
	}

	if price != 0 {
		fmt.Printf("[%v]的年龄为[%v], 门票价格为:%v元。\n", Visitor.Name, Visitor.Age, price)
	} else {
		fmt.Printf("[%v]的年龄为[%v], 门票价格免费。\n", Visitor.Name, Visitor.Age)
	}
}

func (student *Student) Say() string {
	stuInfo := fmt.Sprintf("Student Info: ID=[%v], Name=[%v], Gender=[%v], Age=[%v], Score=[%v]",
		student.Id, student.Name, student.Gender, student.Age, student.Score)
	return stuInfo
}

func main() {
	//var student = Student{
	//	Id:     1,
	//	Name:   "WeiMin",
	//	Gender: "M",
	//	Age:    190,
	//	Score:  120.0,
	//}
	//fmt.Println(student.Say())
	var visitor Visitor
	for {
		visitor.InputData()
		price := visitor.Charge()
		visitor.Print(price)
	}
}
