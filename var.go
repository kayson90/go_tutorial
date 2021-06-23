package main

import "fmt"

var c, python, java bool //包级别

func main() {
	var i int //函数级别
	fmt.Println(i, c, python, java)
}
