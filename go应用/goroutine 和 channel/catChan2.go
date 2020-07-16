//  创建一个catChan2,最多可以存放10个*Cat变量，演示写入和读取的用法
package main

import "fmt"

type Cat struct {
	Name string
	Age  int
}

func main() {
	var catChan chan *Cat
	catChan = make(chan *Cat, 10)
	cat1 := Cat{Name: "tom", Age: 16}
	cat2 := Cat{Name: "tom~", Age: 180}
	catChan <- &cat1
	catChan <- &cat2
	// 取出
	cat11 := <-catChan
	cat22 := <-catChan
	fmt.Println(cat11, cat22)
}

/*

PS D:\goLang\github\golang_project\go应用\goroutine 和 channel> go run .\catChan2.go
				&{tom 16} &{tom~ 180}
*/
