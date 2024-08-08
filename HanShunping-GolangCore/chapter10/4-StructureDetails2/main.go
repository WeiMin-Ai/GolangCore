package main

import (
	"encoding/json"
	"fmt"
)

type Point struct {
	x int
	y int
}

type Rect struct {
	leftUp, rightDown Point
}

type Rect2 struct {
	leftUp, rightDown *Point
}

type Monster struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Skill string `json:"skill"`
}

func main() {
	r1 := Rect{Point{1, 2}, Point{3, 4}}
	fmt.Printf("r1 point is: %p\n", &r1.leftUp.x)
	fmt.Printf("r1 point is: %p\n", &r1.leftUp.y)
	fmt.Printf("r1 point is: %p\n", &r1.rightDown.x)
	fmt.Printf("r1 point is: %p\n\n", &r1.rightDown.y)

	r2 := Rect2{&Point{22, 44}, &Point{55, 66}}
	fmt.Printf("r2 leftUp 本身地址: %p\n", &r2.leftUp)
	fmt.Printf("r2 rightDown 本身地址: %p\n", &r2.rightDown)
	fmt.Printf("r2 point is: %p\n", &r2.leftUp.x)
	fmt.Printf("r2 point is: %p\n", &r2.leftUp.y)
	fmt.Printf("r2 point is: %p\n", &r2.rightDown.x)
	fmt.Printf("r2 point is: %p\n", &r2.rightDown.y)

	monster := Monster{"牛魔王", 5000, "Code!"}
	jsonStr, err := json.Marshal(monster)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(jsonStr))

}
