package main

import "fmt"

var i, j int = 1, 2

func main() {
	var c, python, java = true, false, "no!" //类型推导
	fmt.Println(i, j, c, python, java)
}
