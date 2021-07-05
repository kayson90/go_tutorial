package main

import (
	"fmt"
	"time"
)

//error类型是一个builtin接口
type MyError struct {
	When time.Time
	What string
}

//MyError实现了这个接口
func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s",
		e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}
