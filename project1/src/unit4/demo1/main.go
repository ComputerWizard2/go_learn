package main

import "fmt"

// learn Array by example
func main() {
	var scores [5]int = [5]int{10, 20, 30, 40, 50}
	sum := 0
	for i := 0; i < len(scores); i++ {
		sum += scores[i]
	}
	fmt.Println(sum)
	// 平均值
	avge := sum / len(scores)
	fmt.Println(avge)
}
,