package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])
	//要明白value为0不代表key是存在的，如果要知道是否存在使用下面的方式
	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])
	//取不到是false
	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}
