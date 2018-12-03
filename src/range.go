// range 迭代各种各样的数据结构。slice string arrays map 等
package main

import "fmt"

func main() {

	// 我们使用range 来统计一个 slice 的各位元素的值之和
	// 数组也可以采用这种方法

	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	// 使用range迭代map中的健值对
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Println("%s -> %s\n", k, v)
	}

	// range 在字符串中迭代 unicode 编码。第一个返回值是
	// go 的其实字节位置，然后第二个是 go 自己
	for i, c := range "go" {
		fmt.Println(i, c)
	}
}
