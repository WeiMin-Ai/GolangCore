package main

import (
	"GolangCore/HanShunping-GolangCore/chapter10/11-Factory/model"
	"fmt"
)

func main() {
	var s = *model.NewStudent("WeiMin", 445)
	fmt.Println(s.GetAge())
}
