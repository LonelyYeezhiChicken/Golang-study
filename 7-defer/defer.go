/*
 * @Description: defer关键字
 * @Author: HCQ
 * @Company(School): UCAS
 * @Email: 1756260160@qq.com
 * @Date: 2021-02-14 23:08:07
 * @LastEditTime: 2021-02-16 11:30:42
 * @FilePath: /go-learning/8小时转职Golang工程师/GolangStudy/7-defer/defer.go
 */
package main

import "fmt"

func main() {
	//写入defer关键字
	defer fmt.Println("main end1") // 函数结束前触发执行
	defer fmt.Println("main end2")


	fmt.Println("main::hello go 1")
	fmt.Println("main::hello go 2")
}