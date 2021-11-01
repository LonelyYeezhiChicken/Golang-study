/*
 * @Description:  常量和枚举类型 https://www.bilibili.com/video/BV1gf4y1r79E?p=6
 * @Author: HCQ
 * @Company(School): UCAS
 * @Email: 1756260160@qq.com
 * @Date: 2021-02-14 23:08:07
 * @LastEditTime: 2021-02-15 22:22:13
 * @FilePath: /go-learning/8小时转职Golang工程师/GolangStudy/3-const_iota/test2_const.go
 */
package main

import "fmt"

//const 来定义枚举类型
const (
	//可以在const() 添加一个关键字 iota， 每行的iota都会累加1, 第一行的iota的默认值是0
	BEIJING = 10*iota	 //iota = 0
	SHANGHAI 		  //iota = 1
	SHENZHEN          //iota = 2
)

const (
	a, b = iota+1, iota+2 // iota = 0, a = iota + 1, b = iota + 2, a = 1, b = 2
	c, d				  // iota = 1, c = iota + 1, d = iota + 2, c = 2, d = 3
	e, f				  // iota = 2, e = iota + 1, f = iota + 2, e = 3, f = 4

	g, h = iota * 2, iota *3  // iota = 3, g = iota * 2, h = iota * 3, g = 6, h = 9 
	i, k					   // iota = 4, i = iota * 2, k = iota * 3 , i = 8, k = 12
)

func main() {
	//常量(只读属性)
	const length int = 10

	fmt.Println("length = ", length)

	//length = 100 //常量是不允许修改的。

	fmt.Println("BEIJIGN = ", BEIJING) // BEIJIGN =  0
	fmt.Println("SHANGHAI = ", SHANGHAI) // SHANGHAI =  10
	fmt.Println("SHENZHEN = ", SHENZHEN) // SHENZHEN =  20

	fmt.Println("a = ", a, "b = ", b)
	fmt.Println("c = ", c, "d = ", d)
	fmt.Println("e = ", e, "f = ", f)

	fmt.Println("g = ", g, "h = ", h)
	fmt.Println("i = ", i, "k = ", k)

	// iota 只能够配合const() 一起使用， iota只有在const进行累加效果。
	//var a int = iota 

}