package main

import "fmt"

// 全局变量
var n1 = 100
var n2 = 200
var n3 = 300

// 一次声明多个全局变量
var (
	n7 = 400
	n8 = 500
)

func main() {
	//局部变量
	fmt.Println("Hello, World!")
	// 指定类型并且赋值
	var name string = "John"
	fmt.Println(name)

	// 不指定类型，根据值推断类型
	var num = "10"
	fmt.Println(num)

	// 简短声明
	sex := "male"
	fmt.Println(sex)
	fmt.Printf("=============\n")
	// 声明多个变量
	var a, b, c int
	fmt.Println(a, b, c)
	//声明多个变量，并赋值
	var n4, n5, n6 = 10, "jack", false
	fmt.Println(n4, n5, n6)
	fmt.Printf("n1=%v\n", n1)
	fmt.Printf("n2=%v\n", n2)
	fmt.Printf("n3=%v\n", n3)
	fmt.Printf("n4=%v\n", n7)
	fmt.Printf("n5=%v\n", n8)

}
