package main

import (
	"fmt"
	"sync"
)

/*
channel:1.是一种数据类型,引用类型,需要开辟内存空间(make)后才能使用, var 标识符 chan 元素类型
        2.是goroutine之间数据传递的通道
*/

func main() {
	var (
		chan1 chan int
		wg    sync.WaitGroup
	)

	chan1 = make(chan int) // 初始化不带缓冲区大小chan, 指定缓冲区大小初始化make(chan int, 10)

	wg.Add(1)
	go func() {
		defer wg.Done()
		x := <-chan1
		fmt.Println("get x from chan1:", x)
	}()

	chan1 <- 10  // chan必须有goroutine接收通道值,否则会引起死锁.一直在等待接收的goroutine
	close(chan1) // 关闭通道

	wg.Wait()
}
