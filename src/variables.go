package main

import "fmt"

func main(){

	// var 声明一个或者多个变量
	var a string = "initial"
	fmt.Println(a)

	// 你可以一次性声明多个变量
	var b,c int = 1,2
	fmt.Println(b,c)

	// GO 将自动推断已经初始化的变量类型
	var d = true
	fmt.Println(d)

	// 声明变量且没有给出对应的初始值时，变量将会初始化为
	//例如，一个 int的零值是"0"
	var e int
	fmt.Println(e)

	// ':='语句是声明并初始化变量的简写,例如
	f := "short"
	fmt.Println(f)

}
