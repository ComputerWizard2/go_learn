package main

import (
	"fmt"
)

func main() {
	test()
}

func test() {
	defer func() {
		errors := recover()
		if errors != nil {
			fmt.Println(errors)
		}
	}()
	num1 := 10
	num2 := 0
	num3 := num1 / num2
	fmt.Println(num3)
}
