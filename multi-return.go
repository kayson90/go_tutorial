package main

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}
func main() {
	//声明变量时进行赋值简写
	a, b := swap("hello", "world")
	fmt.Println(a, b)
}
