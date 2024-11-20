package main

import "fmt"

func main() {
	_, ok := add(1, 2)
	fmt.Println(ok)
}
func add(a int, b int) (int, bool) {
	return a + b, true
}
