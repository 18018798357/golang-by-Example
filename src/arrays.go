// Go中数组是一个固定长度的数列
package main

import "fmt"

func main() {

	// 这里我们创建了一个数组'a'来存放5个'int'
	// 元素的类型和长度都是数组的一部分。数组的默认是为对应类型的默认值,如: emp [0 0 0 0 0]
	var a [5]int
	fmt.Println("emp", a)

	// 我们使用'array[index] = value' 语法来设置数组
	// 指定位置的值,或者使用'array[index]'得到值
	a[4] = 100
	fmt.Println("set", a)
	fmt.Println("get", a[4])

	// 使用内置函数'len'返回数组得长度
	fmt.Println("len", len(a))

	// 使用这个语法在一行代码初始化数组
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl", b)

	// 构造多维数组
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}

	fmt.Printf("2d %p", &twoD)
}
