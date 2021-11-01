/*
 * @Description: 基础 https://www.bilibili.com/video/BV1gf4y1r79E?p=4&spm_id_from=pageDriver
 * @Author: HCQ
 * @Company(School): UCAS
 * @Email: 1756260160@qq.com
 * @Date: 2021-02-14 23:08:07
 * @LastEditTime: 2021-02-14 23:41:49
 * @FilePath: /go-learning/8小时转职Golang工程师/GolangStudy/1-firstGolang/hello.go
 */
package main //程序的包名

/*
import "fmt"
import "time"
*/
// 建议导包下面方式
import (
	"fmt"
	"time"
)


//main函数
func main() { //函数的{  一定是 和函数名在同一行的，否则编译错误
	//golang中的表达式，加";", 和不加 都可以，建议是不加
	fmt.Println(" hello Go!")

	time.Sleep(1 * time.Second)
}