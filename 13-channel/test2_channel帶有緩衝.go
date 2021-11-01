package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int, 3) //带有缓冲的channel

	fmt.Println("len(c) = ", len(c), ", cap(c)", cap(c))

	// 发送
	go func() {
		defer fmt.Println("子go程结束")

		for i := 0; i < 4; i++ {
			c <- i   // 写入管道
			fmt.Println("子go程正在运行, 发送的元素=", i, " len(c)=", len(c), ", cap(c)=", cap(c))
		}
	}()

	time.Sleep(2 * time.Second)   // 睡两秒
	// 获取
	for i := 0; i < 4; i++ {
		num := <-c //从c中接收数据，并赋值给num   （<-c  中间没有空格）
		fmt.Println("num = ", num)
	}

	fmt.Println("main 结束")
}
