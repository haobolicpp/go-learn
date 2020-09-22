package main

import (
	"fmt"
	"time"
)

/*
select 语句用于在多个发送/接收信道操作中进行选择。select 语句会一直阻塞，直到发送/接收操作准备就绪。
如果有多个信道操作准备完毕，select 会随机地选取其中之一执行
*/

func main() {
	a := make(chan int)
	b := make(chan int)
	go server1(a)
	go server2(b)
	select {
	case c := <-a:
		fmt.Println(c) //执行该语句
		break
	case d := <-b:
		fmt.Println(d)
		break
		// default: //没有数据则执行此处，防止阻塞
		// 	fmt.Println("no data")
		// 	break
	}
}

func server1(data chan int) {
	time.Sleep(100 * time.Millisecond)
	data <- 1
}

func server2(data chan int) {
	time.Sleep(500 * time.Millisecond)
	data <- 2
}
