package main

import "fmt"

func main() {
    // 创建一个整型带两个缓冲的通道
    ch := make(chan int, 2)
   
    // 给通道放入两个数据
    ch <- 0
    ch <- 1
   
    // 关闭缓冲
    close(ch)

    // 遍历缓冲所有数据, 且多遍历1个
    for i := 0; i < cap(ch)+1; i++ {
   
        // 从通道中取出数据
        v, ok := <-ch
       
        // 打印取出数据的状态
        fmt.Println(v, ok)
    }
}

/**
PS D:\goLang\github\golang_project\Go语言并发\Go语言关闭通道后继续使用通道> go run 2.go
0 true
1 true
0 false

代码说明如下：
第 7 行，创建一个能保存两个元素的带缓冲的通道，类型为整型。
第 10 行和第11行，给这个带缓冲的通道放入两个数据。这时，通道装满了。
第 14 行，关闭通道。此时，带缓冲通道的数据不会被释放，通道也没有消失。
第 17 行，cap() 函数可以获取一个对象的容量，这里获取的是带缓冲通道的容量，也就是这个通道在 make 时的大小。虽然此时这个通道的元素个数和容量都是相同的，但是 cap 取出的并不是元素个数。这里多遍历一个元素，故意造成这个通道的超界访问。
第 20 行，从已关闭的通道中获取数据，取出的数据放在 v 变量中，类型为 int。ok 变量的结果表示数据是否获取成功。
第 23 行，将 v 和 ok 变量打印出来
*/