// 结构体案例演示

package main

import "fmt"

type Cat struct {
	Name  string
	Age   int
	Color string
	Hobby string
}

type Person struct {
	Name   string
	Age    int
	Scores [5]float64
	ptr    *int
	slice  []int
	map1   map[string]string
}

func main() {
	/*
		张老太养猫问题：张老太养了20只猫，每只猫都用名字、年龄、颜色。
		请编写一个程序，当用户输入小猫名字的时候，显示猫咪的名字、年龄、颜色。
		如果用户输入错误，则显示张老太没有这只猫。
	*/

	// 如果没有结构体，如何解决该问题？
	// 1. 变量解决
	/*
		var cat1Name string = "小白"
		var cat1Age int = 3
		var cat1Color string = "白色"
		var cat2Name string = "小黑"
		var cat2Age int = 3
		var cat2Color string = "黑色"
	*/

	// 2. 数组解决
	/*
		var catNames [2]string = [...]string{"小白", "小黑"}
		var catAges [2]int = [...]int{3, 100}
		var catColor [2]string = [...]string{"白色", "黑色"}
	*/

	// 3. map解决
	/*
		var catNames map[string]string = map[string]string{"小黑": "小黑"}
		var catAges map[string]int = map[string]int{"小黑": 3, "小白": 100}
		var catColor map[string]string = map[string]string{"小黑": "黑色", "小白": "白色"}
	*/

	// 4. 结构体解决
	var cat1 Cat
	cat1.Name = "小白"
	cat1.Color = "白色"
	cat1.Age = 3
	cat1.Hobby = "晒太阳"
	fmt.Println("cat1: ", cat1)
	fmt.Println("cat1 Name: ", cat1.Name)
	fmt.Println("cat1 Color: ", cat1.Color)
	fmt.Println("cat1 Age: ", cat1.Age)
	fmt.Println("cat1 Hobby: ", cat1.Hobby)

	// 结构体注意事项：
	// 1. 指针、slice、map的默认值都是nil，即没有分配内存空间。
	var person1 Person
	if person1.ptr == nil {
		fmt.Println("person1.ptr is nil") // nil
	}

	if person1.slice == nil {
		fmt.Println("person1.slice is nil") // nil
	}

	if person1.map1 == nil {
		fmt.Println("person1.map1 is nil") // nil
	}

	// 2. 使用slice一定要使用make
	person1.slice = make([]int, 10)
	person1.slice[0] = 100

	// 3. 使用map一定要使用make
	person1.map1 = make(map[string]string)
	person1.map1["key1"] = "WeiMin"

	fmt.Println(person1)

	// 4. 不同的结构体变量的字段都是独立的，互不影响，因为指向的都是不同的内存地址。
	var cat2 Cat
	cat2.Name = "小黑"
	cat2.Color = "黑色"
	cat2.Age = 3
	cat2.Hobby = "睡觉"
	fmt.Println("cat1: ", cat1)
	fmt.Println("cat2: ", cat2)
}
