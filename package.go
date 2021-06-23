package main

import (
	"fmt"
	"math/rand"
)

func main() {
	//seed是固定的，所以每次生成的随机数是一样的
	fmt.Println("My favourite number is", rand.Intn(10))
}
