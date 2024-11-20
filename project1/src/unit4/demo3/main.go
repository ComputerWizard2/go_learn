package main

import "fmt"

func main() {
	array := [5]int{1, 2, 3, 4, 5}
	//传入
	test(&array)
	fmt.Println(array)
}
func test(array *[5]int) {
	array[2] = 100
}
