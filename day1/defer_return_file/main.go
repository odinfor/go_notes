package main

import "fmt"

// defer和return的执行顺序(return分匿名和显示返回)
// 函数执行顺序
// 1.return前赋值
// 2.执行defer
// 3.执行RET

// 1.匿名返回值,x:=5给返回值赋值. 2.执行defer改变的是原x不是返回值的赋值. 3.RET x=5
func f1() int {
	x := 5
	defer func() {
		x++
	}()
	return x
}

// 显示返回 1.return 5 给返回值x赋值5. 2.执行defer改变x的值. 3.RET x=6
func f2() (x int) {
	defer func() {
		x++
	}()
	return 5
}

// 显示返回 1.x:=5 return x 给y赋值5. 2.执行defer改变x的值. 3.RET y=5
func f3() (y int) {
	x := 5
	defer func() {
		x++
	}()
	return x
}

// 显示返回 1.return 5 给x赋值5. 2.执行defer将x作为参数传入,命名空间作用域为x的副本. 3.RET x=5
func f4() (x int) {
	defer func(x int) {
		x++
	}(x)
	return 5
}

func main() {
	fmt.Println(f1())
	fmt.Println(f2())
	fmt.Println(f3())
	fmt.Println(f4())
}
