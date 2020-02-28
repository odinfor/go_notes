package main

import "fmt"

// 切片相关,切片的容量和切片与底层数组的关系.
// 切片的本质：一块连续的内存.故只能存放相同的数据类型.

func main() {
	var array1 = [10]int{0: 1, 3: 4, 4: 3, 6: 9, 8: 11} // 创建数组 [1 0 0 4 3 0 9 0 11 0]

	// 创建切片 1：初始化创建: 2：由底层数组创建
	slice1 := []int{1, 2, 3, 4, 5, 6, 7} // 初始化创建
	slice2 := array1[:4]                 // 由底层数组创建
	slice3 := array1[2:4]

	fmt.Println("array1: ", array1)

	// 打印切片和切片的容量
	// 切片的容量为切片后底层数组的最大可扩展长度.从左至右,不包含左边起始index前的.
	// 初始化创建的切片实则为两步.1.创建一个数组, 2.创建数组[:]的切片.
	fmt.Println("slice1: ", slice1, "cap: ", cap(slice1))
	fmt.Println("slice2: ", slice2, "cap: ", cap(slice2))
	fmt.Println("slice3: ", slice3, "cap: ", cap(slice3)) // [0 4], 容量还能扩展[3 0 9 0 11 0] 故cap=8

	// 由于切片属引用传递,切片没有自己的值,而是指向其底层数组,故修改底层数组的值,切片的值将一并改变.
	array1[3] = 10
	fmt.Println("slice2: ", slice2, "cap: ", cap(slice2))
	fmt.Println("slice3: ", slice3, "cap: ", cap(slice3))

	// 1.切片本身没有删除元素操作方法,需要借助append来实现
	// 2。切片删除元素操作: 将会改变底层数组
	a1 := [...]int{1, 2, 3, 4, 5, 6, 7}
	s1 := a1[:] // 生成切片

	fmt.Printf("%p\n", &a1[0])     // 修改前数组a1 index=0的元素的地址
	s1 = append(s1[:2], s1[3:]...) // 将index为2的元素移除,得到[1 2 4 5 6 7]的切片.此时数组a1发生来改变.
	fmt.Printf("%p\n", &a1[0])     // 修改后数组a1 index=0的元素的地址

	// 由上可知,数组a1的元素地址没有改变.即append并没有产生数组copy操作.
	// 切片本身不存值,是底层数组的引用.append改变切片的元素其实是对底层数组的修改.

	fmt.Println(s1, len(s1), cap(s1))
	fmt.Println(a1) // a1=[1 2 4 5 6 7 7]
}
