// 函数是go的中心
package main

import "fmt"

func plus(a int, b int) int {

	//Go 需要明确的返回值，例如它不会自动但会最后一个表达是的值
	return a + b
}

func main() {
	// 正如你期望的那样，通过 name(args)来调用一个函数
	res := plus(1, 2)
	fmt.Println("1+2=", res)
}
