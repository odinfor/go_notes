package main

import "fmt"

// 指针,创建指针
// new&make均用于申请内存
// new:用与申请内存-用于基本数据类型(值类型)：int/string... 返回该数据类型的指针,*int/*string...
// make:用与内存分配-只用与:slice/map/chan(引用类型),返回类型本身

func main() {
	// 创建一个指针
	var p1 *int // 该方法创建的指针,没有内存地址,为nil,不能为指针赋值.
	fmt.Println(p1)
	// *p1 = 100		// panic: runtime error: invalid memory address or nil pointer dereference

	var p2 = new(int) // 该方法创建的指针,具有内存地址,不为nil,可以为指针赋值.
	fmt.Println(p2)
	fmt.Println(*p2) // 0
	*p2 = 100
	fmt.Println(*p2)
}
