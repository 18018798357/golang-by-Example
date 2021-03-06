/**
	if 和 else 分之结构在GO中当然是直接了当的
 */
package main

import "fmt"

func main() {

	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	// 你可以不要else 只用 if语句
	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}

	// 在条件语句之前可以有一个语句;任何在这里声明的变量都可以在所有的条件分支中使用
	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}

	// 注意,在Go中你可以不使用圆括号,但是花括号是需要的
}
