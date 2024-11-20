package main

import "fmt"

func main() {
	a := 0
	sum := 0
	// for 循环
	for i := 0; i < 10; i++ {
		sum += a
		a++
	}
	fmt.Println(sum)

	str := "hello"
	for i, v := range str {
		fmt.Printf("index: %d, value: %c\n", i, v)
	}
}
