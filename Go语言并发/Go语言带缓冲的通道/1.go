package main 

import (
	"fmt"
)

func main(){
	// 创建一个3个元素缓冲大小的整型通道
	ch:=make(chan int,3) // 缓冲大小：决定通道最多可以保存的元素数量。
	// 查看当前通道的大小
	fmt.Println(len(ch))

	// 发送3个整型元素到通道
	ch <- 1
	ch <- 2
	ch <- 3 
	// 查看当前通道的大小
	fmt.Println(len(ch))
}

/**
PS D:\goLang\github\golang_project\Go语言并发\Go语言带缓冲的通道> go run 1.go
0
3

代码说明如下：
第 8 行，创建一个带有 3 个元素缓冲大小的整型类型的通道。
第 11 行，查看当前通道的大小。带缓冲的通道在创建完成时，内部的元素是空的，因此使用 len() 获取到的返回值为 0。
第 14～16 行，发送 3 个整型元素到通道。因为使用了缓冲通道。即便没有 goroutine 接收，发送者也不会发生阻塞。
第 19 行，由于填充了 3 个通道，此时的通道长度变为 3。
*/