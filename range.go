package main

import "fmt"

var powArray = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	for i, v := range powArray {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}
