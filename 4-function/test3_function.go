/*
 * @Description: Golang中函数的多返回值三种写法   https://www.bilibili.com/video/BV1gf4y1r79E?p=7
 * @Author: HCQ
 * @Company(School): UCAS
 * @Email: 1756260160@qq.com
 * @Date: 2021-02-14 23:08:07
 * @LastEditTime: 2021-02-15 22:34:25
 * @FilePath: /go-learning/8小时转职Golang工程师/GolangStudy/4-function/test3_function.go
 */
 package main

 import "fmt"
 
 func foo1(a string, b int) int {
	 fmt.Println("a = ", a)
	 fmt.Println("b = ", b)
 
	 c := 100
 
	 return c
 }
 
 //返回多个返回值，匿名的
 func foo2(a string, b int) (int, int) {
	 fmt.Println("a = ", a)
	 fmt.Println("b = ", b)
 
	 return 666, 777
 }
 
 //返回多个返回值， 有形参名称的
 func foo3(a string, b int) (r1 int, r2 int) {
	 fmt.Println("---- foo3 ----")
	 fmt.Println("a = ", a)
	 fmt.Println("b = ", b)
 
 
 
	 //r1 r2 属于foo3的形参，  初始化默认的值是0
	 //r1 r2 作用域空间 是foo3 整个函数体的{}空间
	 fmt.Println("r1 = ", r1)
	 fmt.Println("r2 = ", r2)
 
 
	 //给有名称的返回值变量赋值
	 r1 = 1000
	 r2 = 2000
 
	 return
 }
 
 func foo4(a string, b int) (r1, r2 int) {
	 fmt.Println("---- foo4 ----")
	 fmt.Println("a = ", a)
	 fmt.Println("b = ", b)
 
 
	 //给有名称的返回值变量赋值
	 r1 = 1000
	 r2 = 2000
 
	 return
 }
 
 func main() {
	 c := foo1("abc", 555)
	 fmt.Println("c = ", c)
 
	 ret1, ret2 := foo2("haha", 999)
	 fmt.Println("ret1 = ", ret1, " ret2 = ", ret2)
 
	 ret1, ret2 = foo3("foo3", 333)
	 fmt.Println("ret1 = ", ret1, " ret2 = ", ret2)
 
	 ret1, ret2 = foo4("foo4", 444)
	 fmt.Println("ret1 = ", ret1, " ret2 = ", ret2)
 }