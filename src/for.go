/**
	for是Go中唯一的循环结构。这里有for循环的三个基本使用方式
 */
package main

import "fmt"

func main() {

	// 最常用的方式,带单个循环条件
	i := 1
	for i < 3 {
		fmt.Println(i)
		i++
	}

	// 经典的初始化/条件/后续形式for循环
	for j := 7; j <= 20; j++ {
		fmt.Println(j)
	}

	// 不带条件的for循环一直执行,直到在循环体内使用了 break 或者 return 来跳出循环
	k := 1
	for {
		fmt.Println("loop", k)
		k++

	}
}
