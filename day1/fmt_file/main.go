package main

import "fmt"

// fmt.Printf 占位符

func main() {
	var n = 100
	var s = "this is s"

	fmt.Printf("%T\n", n) // %T:打印变量类型
	fmt.Printf("%v\n", n) // %v:打印变量值
	fmt.Printf("%b\n", n) // %d:打印变量2进制值
	fmt.Printf("%d\n", n) // %d:打印变量10进制值
	fmt.Printf("%o\n", n) // %o:打印变量8进制值
	fmt.Printf("%x\n", n) // %x:打印变量16进制值

	fmt.Printf("%s\n", s) // %s:打印字符串变量
}
