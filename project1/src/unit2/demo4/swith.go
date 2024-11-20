package main

import "fmt"

func main() {
	var score  := 90
	// switch 语句 是一个表达式，变量 常量
	switch score / 10 {
	case 9:
		fmt.Println("优秀")
		// 穿透
		fallthrough
	case 8:
		fmt.Println("良好")
	case 7:
		fmt.Println("中等")
	case 6:
		fmt.Println("及格")
	default:
		fmt.Println("不及格")
	}
}
