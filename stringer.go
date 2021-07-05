package main

import "fmt"

type Person2 struct {
	Name string
	Age  int
}

//实现的是fmt里面的String()方法
func (p Person2) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main() {
	a := Person2{"Arthur Dent", 42}
	z := Person2{"Zaphod Beeblebrox", 9001}
	fmt.Println(a, z)
}
