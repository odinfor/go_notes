package main

/*
值和地址的区别
*/

// 地址交换
func swap(a, b *int) (*int, *int) {
	a, b = b, a
	return a, b
}

func main() {
	a, b := 3, 4
	c, d := swap(&a, &b) // c是a的地址: {xxxx} 4; d是b的地址: {xxxx} 3
	println(c, d)
	println(*c, *d) // 4  3
	a = *c          // 修改a的值,将a改为c地址指向的值,c={xxxx} 4,c的值为*c=4

	// 上一步执行完之后a,b的值变为了4,4

	b = *d // 修改b的值,a,b此时值一样,地址改变后指向的值均为4
	println(&a, &b)
	println(a, b) // 4  4

	for i, y := 0, 0; i < 10 && y < 5; i++ {

	}
}
