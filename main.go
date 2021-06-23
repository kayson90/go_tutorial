package main

import "fmt"

type Person struct {
	name string
	age  int
}

func main() {
	fmt.Println("Hello go!")
	fmt.Println(f())

	var person = Person{"zhangsan", 8}
	fmt.Println(person)

	setName(person, "lisi")
	fmt.Println(person)

	setName2(&person, "lisi")
	fmt.Println(person)
	//开启了一个goroutine
	go f()
}

func f() int {
	return 8
}

func setName(person Person, name string) {
	person.name = name
}

func setName2(person *Person, name string) {
	(*person).name = name
}
