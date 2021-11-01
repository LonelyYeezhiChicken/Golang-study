/*
 * @Description: 切片
 * @Author: HCQ
 * @Company(School): UCAS
 * @Email: 1756260160@qq.com
 * @Date: 2021-02-14 23:08:07
 * @LastEditTime: 2021-02-16 11:39:00
 * @FilePath: /go-learning/8小时转职Golang工程师/GolangStudy/8-slice/test2_slice.go
 */
package main

import "fmt"

func printArray(myArray []int) {
	//引用传递
	// _ 表示匿名的变量
	for _, value := range myArray {
		fmt.Println("value = ", value)
	}

	myArray[0] = 100
}

func main() {
	myArray := []int{1,2,3,4} // 动态数组，切片 slice（[]中间没有数字）

	fmt.Printf("myArray type is %T\n", myArray)   // myArray type is []int

	printArray(myArray)  // 引用

	fmt.Println(" ==== ")

	for _, value := range myArray {
		fmt.Println("value = ", value)
	}
}