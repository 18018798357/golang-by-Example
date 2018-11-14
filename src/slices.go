



/**
	slice 是Go中一个关键的数据类型,是一个比数组更加强大的序列接口
 */
package main

import "fmt"

func main(){

	// 不像数组,slice的类型仅有它包含的元素决定(不像数组中还需要元素的个数).
	// 要创建一个床都非零的空clise,需要使用make
	// 这里我们创建了一个长度为3的string 类型slice(初始化为零值)
	s:=make([]string,3)
	fmt.Println("emp",s)

	// 我们可以和数组一起设置得到值
	s[0] ="a"
	s[1] ="b"
	s[2] ="c"
	fmt.Println("set",s)
	fmt.Println("get",s[2])
	// 如你所料,'len'返回 slice的长度
	fmt.Println("len:",len(s))

	// 作为基本操作的补充,slice 支持比数组更多的操作
	// 其中一个内建 'append',它返回一个包含了一个或者多个新值的slice.
	// 注意我们接受的返回新的slice值
	s = append(s,"d")
	s = append(s,"e","f")
	fmt.Println("apd",s)

	// slice 支持通过'slice[low:high]'语法进行切片操作.例如,这里得到一个包含元素's[2]','s[3]'
	// 's[4]'的slice
	l:= s[2:5]
	fmt.Println(l)

	//这个slice从 's[0]'到(但是不包含)'s[5]'
	l = s[:5]
	fmt.Println("sl2",l)

	// 这个slice从(包含)'s[2]'到slice的后一个值
	l=s[2:]
	fmt.Println("sl3",l)

	// 我们可以在一行代码中声明并初始化一个slice量
	t:=[]string{"g","h","i"}
	fmt.Println("dcl:",t)

	// slice可以组成多维数组结构。内部的slice长度可以不同,这和多位数组不同
	twoD := make([][]int,3)
	for i:=0;i<3 ;i++  {
		innerLen := i+1
		twoD[i] = make([]int,innerLen)
		for j:=0;j<innerLen ;j++  {
			twoD[i][j]=i+j
		}
	}

	fmt.Println("2d",twoD)
}
