package main

import "fmt"

func run(arg string) {
	fmt.Println(arg)
}

func main() {
	go run("test")
}
