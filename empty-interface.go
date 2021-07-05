package main

import "fmt"

//空接口非常有用
//空接口被用来处理未知类型的值,
func main() {
	var i interface{}
	describe4(i)

	i = 42
	describe4(i)

	i = "hello"
	describe4(i)
}

func describe4(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}
