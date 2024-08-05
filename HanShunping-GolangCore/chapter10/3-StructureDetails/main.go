package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	var p1 Person
	p1.Name = "WeiMin"
	p1.Age = 9999

	// 注意了！
	var p2 *Person = &p1
	fmt.Println(p1, p2)
	p2.Name = "Tom"
	fmt.Println(p1, p2)

	// 内存地址
	fmt.Printf("p1的内存地址：%p\n", &p1)
	fmt.Printf("p2的内存地址：%p, p2字段值的内存地址: %p", &p2, p2)
}
