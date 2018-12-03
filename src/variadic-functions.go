// 数量的参数调整
package main

import "fmt"

func sum(nums ...int) {

	fmt.Println(nums, "")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)

}

func main() {
	// 变参函数使用常规的调用方式,出了参数比较特殊
	sum(1, 2)
	sum(1, 2, 3)

	//	如果你的slice 已经有了多个值,想把它们作为变参使用，你要这样调用 func(slice...)
	nums := []int{1, 2, 3, 4}
	sum(nums...)

}
