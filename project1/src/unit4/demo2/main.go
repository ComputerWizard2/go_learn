package main

import "fmt"

func main() {
	array := [3]int{1, 3, 4}
	var array2 [3]int = [3]int{1, 2, 3}
	fmt.Println(array, array2)
	for i := 0; i < len(array); i++ {
		array[i] += 1
		fmt.Println(&array[i])
	}
	for i, v := range array2 {
		fmt.Printf("for range %d -- %d \n", i, v)
	}
	fmt.Println(array)

}
