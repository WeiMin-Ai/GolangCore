// 1. 结构体的声明方式

package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

// StructureUsage 结构体的使用方式
func StructureUsage() {
	// 1. 方式1
	var weimin Person

	// 2. 方式2
	var weimin2 Person = Person{"Weimin", 999}
	weimin3 := Person{"Weimin", 999}
	fmt.Println(weimin, weimin2, weimin3)

	// 3. 方式3，两种赋值方式都是正确的。
	var person *Person = new(Person)
	(*person).Name = "WeiMin"
	person.Name = "WeiMin1"
	fmt.Println(person)

	// 4. 方式4
	var person2 *Person = &Person{}
	(*person2).Name = "WeiMin2"
	person2.Name = "WeiMin3"
	fmt.Println(person2)

	/*
			说明：
			1.方式3、方式4返回的是结构体指针。
			2.结构体指针访问字段的标准方式应该是：(*结构体指针).字段名
			3.虽然结构体指针的标准方式是第二点，但是go也做了简化，支持结构体指针.字段名进行访问，如person.Name = "Tom\n
		    这么使用也没毛病，go编译器底层还是对person.Name做了转化(*person).Name

	*/
}

func main() {
	// 1. 结构体的声明方式
	StructureUsage()
}
