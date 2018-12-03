// Go 内建多个返回值指出。这个特性在go语言中经常被用到
// 例如用来同时返回一个函数的结果和错误信息

package main

import "fmt"

func vals() (int, int) {

	return 3, 7
}

func main() {

	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	//如果你仅仅想返回值的一部分的话，你可以使用空白定义符'_'
	_, c := vals()
	fmt.Println(c)

}
