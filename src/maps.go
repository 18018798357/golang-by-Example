// map 是Go内置[关联数据类型]
// 在一些其他语言中成为哈希或者字典
package main

import "fmt"

func main() {

	// 要创建一个空map,需要使用内建的'make':
	// make(map [key-type]val-type)
	m := make(map[string]int)

	// 使用典型的'make[key] = val' 语法来设置健值对
	m["k1"] = 7
	m["k2"] = 13

	//输出所有的健值对
	fmt.Println("map:", m)

	prs := m["k2"]
	fmt.Println("prs:", prs)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)
}
