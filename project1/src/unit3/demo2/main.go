package main

import "fmt"

// 交换两个变量
func swap(a int, b int) (int, int) {
	var temp int
	temp = a
	a = b
	b = temp
	return a, b
}

func main() {
	a, b := swap(1, 2)
	fmt.Println(a, b)
}
