package main

import "fmt"

/*
接口的实现,需要实现接口下的所有方法才能称为实现了接口
*/

// 值接收者实现的接口
type speaker interface {
	speak()
	move()
}

type cat struct {
	name string
}

func (c cat) speak() {
	fmt.Printf("%s:miao~", c.name)
}

func (c cat) move() {
	fmt.Printf("%s:zou~", c.name)
}

// 指针接收者实现的接口
type dog struct {
	name string
}

func (d *dog) speak() {
	fmt.Printf("%s:miao~", d.name)
}

func (d *dog) move() {
	fmt.Printf("%s:zou~", d.name)
}

func main() {
	var in1 speaker

	ss := cat{"odin"}
	in1 = ss // cat实现的是值类型接口,赋值需要用值类型
	in1.move()
	in1.speak()

	dog1 := dog{"dog001"}
	in1 = &dog1 // dog实现的是指针类型接口,赋值需要指针类型
	in1.speak()
	in1.move()

}
