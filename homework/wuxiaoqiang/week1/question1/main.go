package main

import (
	"fmt"
	"time"
)

// 两个协程交替打印数字1-100

func main() {
	fmt.Println("beginning of main")
	channel := make(chan int)

	go goroutine1(channel)
	go goroutine2(channel)

	time.Sleep(1 * time.Second)
	fmt.Println("end of main")
}

func goroutine1(ch chan int) {
	for i := 1; i <= 100; i++ {
		ch <- i
		if i%2 == 1 {
			fmt.Println("协程1：", i)
		}
	}
}
func goroutine2(ch chan int) {
	for i := 1; i <= 100; i++ {
		<-ch
		if i%2 == 0 {
			fmt.Println("协程2：", i)
		}
	}
}
