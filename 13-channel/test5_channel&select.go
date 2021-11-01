package main

import "fmt"
// 斐波那契数列  监控两个channel  c&quit
func fibonacii(c, quit chan int) {
	x, y := 1, 1

	for {
		select {
		case c <- x:
			//如果c可写，则该case就会进来
			x = y
			y = x + y // 斐波那契数列
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)

	//sub go
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}

		quit <- 0
	}()

	//main go
	fibonacii(c, quit)
}
