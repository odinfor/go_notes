package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// 通知f1方法结束
// 方式一: 通过修改全局变量判断for循环退出
// 方式二: 通过修改chan值判断for循环退出
// 方式三: 通过context判断for循环退出

var (
	exist  bool = false
	wg     sync.WaitGroup
	exChan = make(chan bool, 1)
	ctx    context.Context
	exCtx  context.CancelFunc
)

// 方式一: 通过修改全局变量判断for循环退出
func f1() {
	for {
		fmt.Println("this is f1")
		time.Sleep(time.Second)
		if exist {
			wg.Done()
		}
	}
}

// 方式二: 通过修改chan值判断for循环退出
func f2() {
	defer wg.Done()
OUT:
	for {
		fmt.Println("this is f2")
		time.Sleep(time.Second)
		select {
		case <-exChan: // 从通道取值
			break OUT // 跳出OUT标签
		default:
		}
	}
}

// 方式三: 使用context取消
// context可向任何子goroutine中传递
func f3(ctx context.Context) {
	defer wg.Done()
OUT:
	for {
		fmt.Println("this is f3")
		time.Sleep(time.Second)
		select {
		// 原理同方式二,只是将chan通过context的Done方法封装了一层用来达到编码风格统一和安全调用
		// 单向通道
		case <-ctx.Done():
			break OUT
		default:
		}
	}
}

// 子goroutine可以有多个子goroutine,通过context实现统一取消
func f4(ctx context.Context) {
	defer wg.Done()
	go f3(ctx)
OUT:
	for {
		fmt.Println("this is f4")
		time.Sleep(time.Second)
		select {
		// 原理同方式二,只是将chan通过context的Done方法封装了一层用来达到编码风格统一和安全调用
		// 单向通道
		case <-ctx.Done():
			break OUT
		default:
		}
	}
}

func main() {
	wg.Add(1)
	// 方式一
	//go f1()
	//time.Sleep(time.Second * 5)
	//exist = true

	// 方式二
	//go f2()
	//time.Sleep(time.Second * 5)
	//exChan <- true

	// 方式三
	//ctx, exCtx = context.WithCancel(context.Background())
	//go f3(ctx)
	//time.Sleep(time.Second * 5)
	//exCtx()

	// 多个子goroutine
	ctx, exCtx = context.WithCancel(context.Background())
	go f4(ctx)
	time.Sleep(time.Second * 5)
	exCtx()

	wg.Wait()
}
