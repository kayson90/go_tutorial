package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

//参数类型一样，可以只写一个类型
func add1(x, y int) int {
	return x + y
}
func main() {
	fmt.Println(add(42, 13))
	fmt.Println(add1(1, 2))
}
