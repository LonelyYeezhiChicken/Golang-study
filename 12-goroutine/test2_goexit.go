package main

import (
	"fmt"
	"time"
)

func main() {

	/*
		//用go创建承载一个形参为空，返回值为空的一个函数
		go func() {  // 匿名函数
			defer fmt.Println("A.defer")

			func() {
				defer fmt.Println("B.defer")
				//退出当前goroutine
				runtime.Goexit()
				fmt.Println("B")
			}()  // 不用声明，自运行

			fmt.Println("A")
		}()   // 不用声明，自运行

	go func(a int, b int) bool {
		fmt.Println("a = ", a, ", b = ", b)
		return true
	}(10, 20)

	//死循环
	for {
		time.Sleep(1 * time.Second)
	}

}
