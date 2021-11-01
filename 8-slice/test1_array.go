/*
 * @Description: 
 * @Author: HCQ
 * @Company(School): UCAS
 * @Email: 1756260160@qq.com
 * @Date: 2021-02-14 23:08:07
 * @LastEditTime: 2021-02-16 11:33:32
 * @FilePath: /go-learning/8小时转职Golang工程师/GolangStudy/8-slice/test1_array.go
 */
package main

import "fmt"

func printArray(myArray [4]int) {
	//值拷贝

	for index, value := range myArray {
		fmt.Println("index = ", index, ", value = ", value)
	}

	myArray[0] = 111
}


func main() {
	//固定长度的数组
	var myArray1 [10]int

	myArray2 := [10]int{1,2,3,4}
	myArray3 := [4]int{11,22,33,44}

	//for i := 0; i < 10; i++ {
	for i := 0; i < len(myArray1); i++ {
		fmt.Println(myArray1[i])	// 默认值为0
	}

	for index, value := range myArray2 {
		fmt.Println("index = ", index, ", value = ", value)
	}

	//查看数组的数据类型
	fmt.Printf("myArray1 types = %T\n", myArray1)
	fmt.Printf("myArray2 types = %T\n", myArray2)
	fmt.Printf("myArray3 types = %T\n", myArray3)

	printArray(myArray3)
	fmt.Println(" ------ ")
	for index, value := range myArray3 {
		fmt.Println("index = ", index, ", value = ", value)
	}
}