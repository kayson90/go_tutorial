package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}

	s = s[1:4]
	fmt.Println(s)
	//对切片操作里边的数组发生了变化
	s = s[:2]
	fmt.Println(s)

	s = s[1:]
	fmt.Println(s)
}
