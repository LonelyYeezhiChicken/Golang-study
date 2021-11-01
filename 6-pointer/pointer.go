/*
 * @Description: 指针
 * @Author: HCQ
 * @Company(School): UCAS
 * @Email: 1756260160@qq.com
 * @Date: 2021-02-14 23:08:07
 * @LastEditTime: 2021-02-16 11:24:03
 * @FilePath: /go-learning/8小时转职Golang工程师/GolangStudy/6-pointer/pointer.go
 */
package main

import "fmt"

/*
func swap(a int ,b int) {
	var temp int
	temp = a
	a = b
	b = temp
}
*/

func swap(pa *int, pb *int) {
	var temp int
	temp = *pa //temp = main::a
	*pa = *pb  // main::a = main::b
	*pb = temp // main::b = temp
}


func main() {
	var a int = 10
	var b int = 20

	swap(&a, &b)

	fmt.Println("a = ", a, " b = ", b)


	var p *int

	p = &a

	fmt.Println(&a)
	fmt.Println(p)

	// 二级指针！！！！！！！！！！！！！！！！！！
	var pp **int //二级指针

	pp = &p

	fmt.Println(&p)
	fmt.Println(pp)
}