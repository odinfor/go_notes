package main

import "fmt"

// 常量创建后不可修改
// iota只在const中,初始化为0,多一行自增+1
const (
	a1 = iota // 0
	a2 = 100
	a3
	a4 = iota
)

const (
	t1 = 10
	t2 = iota // 0
	t3        // 1
)

// iota多值赋值
const (
	d1, d2 = iota, iota + 1     // d1=0, d2=1
	d3, d4 = iota + 1, iota + 2 // d3=2, d4=3
)

func main() {
	fmt.Println("a1:", a1)
	fmt.Println("a2:", a2)
	fmt.Println("a3:", a3)
	fmt.Println("a4:", a4)

	fmt.Println("d1:", d1)
	fmt.Println("d2:", d2)
	fmt.Println("d3:", d3)
	fmt.Println("d4:", d4)
}
