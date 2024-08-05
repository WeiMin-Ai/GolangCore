package main

import "fmt"

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

func main() {
	r1 := Rect{Point{1, 2}, Point{3, 4}}
	fmt.Printf("r1 point is: %p\n", &r1.leftUp.x)
	fmt.Printf("r1 point is: %p\n", &r1.leftUp.y)
	fmt.Printf("r1 point is: %p\n", &r1.rightDown.x)
	fmt.Printf("r1 point is: %p\n", &r1.rightDown.y)
}
