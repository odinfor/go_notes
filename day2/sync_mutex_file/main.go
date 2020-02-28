package main

import (
	"fmt"
	"sync"
	"time"
)

/*
sync.Mutex: 互斥锁
*/

var (
	wg     sync.WaitGroup
	lock   sync.Mutex   // 普通互斥锁
	rwLock sync.RWMutex // 读写互斥锁,读操作远大于写操作时使用,性能远大于互斥锁

	x = 0
)

func f1() {
	for i := 0; i < 1000; i++ {
		lock.Lock() // 对公共变量加锁,保证同一时间对x的操作只有一个
		x += 1
		lock.Unlock() // 释放锁
	}
	wg.Done()
}

// 使用互斥锁模拟读操作,假定每次读取时间耗时1ms
func readWithLock() {
	lock.Lock()
	time.Sleep(time.Microsecond)
	fmt.Println(x)
	lock.Unlock()

	defer wg.Done()
}

// 使用读写互斥锁模拟读操作
func readWIthRwLock() {
	rwLock.RLock()
	time.Sleep(time.Microsecond)
	fmt.Println(x)
	rwLock.RUnlock()

	defer wg.Done()
}

// 使用互斥锁模拟写操作,假定每次写耗时5ms
func writeWithLock() {
	lock.Lock()
	time.Sleep(time.Microsecond * 5)
	x += 5
	lock.Unlock()

	defer wg.Done()
}

func writeWithRwLock() {
	rwLock.RLock()
	time.Sleep(time.Microsecond * 5)
	x += 5
	rwLock.RUnlock()

	defer wg.Done()
}

func main() {
	//wg.Add(2)
	//go f1()
	//go f1()
	//wg.Wait()
	//
	//fmt.Println(x)

	startTime := time.Now()

	// 使用互斥锁模拟
	for i := 0; i < 5000; i++ {
		wg.Add(1)
		go readWithLock()
	}
	// 模拟写10次
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go writeWithLock()
	}

	wg.Wait()

	lockTime := time.Now().Sub(startTime)
	startTime1 := time.Now()

	// 使用读写互斥锁模拟
	for i := 0; i < 5000; i++ {
		wg.Add(1)
		go readWIthRwLock()
	}
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go writeWithRwLock()
	}
	wg.Wait()
	// 使用读写互斥锁比使用互斥锁性能优越太多
	fmt.Println(lockTime)                   // 2s左右
	fmt.Println(time.Now().Sub(startTime1)) // 30~40ms左右
}
