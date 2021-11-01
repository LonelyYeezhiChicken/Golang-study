/*
 * @Description: 拷贝
 * @Author: HCQ
 * @Company(School): UCAS
 * @Email: 1756260160@qq.com
 * @Date: 2021-02-14 23:08:07
 * @LastEditTime: 2021-02-16 11:54:55
 * @FilePath: /go-learning/8小时转职Golang工程师/GolangStudy/8-slice/test5_slice.go
 */
package main

import "fmt"

func main() {
	s := []int{1, 2, 3} //len = 3, cap = 3, [1,2,3]

	//[0, 2)
	s1 := s[0:2] // [1, 2]

	fmt.Println(s1)

	s1[0] = 100

	fmt.Println(s)
	fmt.Println(s1)

	//copy 可以将底层数组的slice一起进行拷贝
	s2 := make([]int, 3) //s2 = [0,0,0]

	//将s中的值 依次拷贝到s2中
	copy(s2, s)
	fmt.Println(s2) // [100 2 3]

}
