package main

import "fmt"

func main() {
	var count int = 100
	if count > 0 {
		fmt.Println("count is greater than 0")
	} else if count < 0 {
		fmt.Println("count is less than 0")
	} else {
		fmt.Println("count is 0")
	}
	// 变量定义
	if x := 10; x > 0 {
		fmt.Println("x is greater than 0")
	}
}
