package main

import "fmt"

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[1:4] //中括号里面是空的就是切片
	fmt.Println(s)
}
