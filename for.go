package main

import "fmt"

func main() {
	sum := 0
	//不需要小括号倒是蛮别致的
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Print(sum)
}
