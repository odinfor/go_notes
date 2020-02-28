package main

import (
	"fmt"
	"sync"
	"time"
)

/*
goroutine&sync.WaitGroup
1.goroutine:创建协程
2.sync.WaitGroup:优雅关闭,等待协程执行完毕后关闭

执行:
1. 添加sync.WaitGroup.Add() 每一个goroutine对应一个sync.WaitGroup
2. sync.WaitGroup.Wait()
3. sync.Done()
*/
var wg sync.WaitGroup

func f1() {
	defer wg.Done() // 释放goroutine
	fmt.Println("start f1")
	time.Sleep(time.Second * 5)
	fmt.Println("f1 done")
}

func f2() {
	defer wg.Done() // 释放goroutine
	fmt.Println("start f2")
	time.Sleep(time.Second * 3)
	fmt.Println("f2 done")
}

func main() {
	//wg.Add(2)
	//go f1()
	//go f2()
	//wg.Wait()	// 等待goroutine执行

	for i := 0; i <= 5; i++ {
		wg.Add(2)
		go f1()
		go f2()
	}
	wg.Wait()
}
