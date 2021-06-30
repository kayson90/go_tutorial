package main

import "fmt"

/**
切片的长度就是它所包含的元素个数
切片的容量是从它的第一个元素开始数，到其底层数组元素末尾的个数。
要注意什么时候切片的容量会发生变化
*/
func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)
	// 截取切片使其长度为 0
	s = s[:0]
	printSlice(s)

	// 拓展其长度,数组的元素会出现
	s = s[:4]
	printSlice(s)
	// 舍弃前两个值，数组的元素不会出现
	s = s[2:]
	printSlice(s)
	s = s[0:4] //拓展长度，数组内的元素出现
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
