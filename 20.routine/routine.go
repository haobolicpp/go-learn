package main

import (
	"fmt"
	"time"
	"sync"
)

/*
调用函数或者方法时，在前面加上关键字 go，可以让一个新的 Go 协程并发地运行
*/

func main() {
	//简单启用
	go hello()
	fmt.Println("main func") //主函数会运行在一个特殊的协程上 Main Goroutine

	/*
		启动一个新的协程时，协程的调用会立即返回。与函数不同，程序控制不会去等待 Go 协程执行完毕。在调用 Go 协程之后，程序控制会立即返回到代码的下一行，忽略该协程的任何返回值。
		如果希望运行其他 Go 协程，Go 主协程必须继续运行着。如果 Go 主协程终止，则程序终止，于是其他 Go 协程也不会继续运行。
	*/

	//休眠一会，阻塞主协程，让hello协程有时间执行
	time.Sleep(100 * time.Millisecond)
	fmt.Println("main func 2")

	/*
		【信道定义】
		信道可以想像成 Go 协程之间通信的管道。如同管道中的水会从一端流到另一端，通过使用信道，数据也可以从一端发送，在另一端接收
		用 make 来定义信道
		【信道读写】
		data := <- a // 读取信道 a
		a <- data // 写入信道 a
		【收发阻塞】
		发送与接收默认是阻塞的。这是什么意思？当把数据发送到信道时，
		程序控制会在发送数据的语句处发生阻塞，直到有其它 Go 协程从信道读取到数据，才会解除阻塞。
		与此类似，当读取信道的数据时，如果没有其它的协程把数据写入到这个信道，那么读取过程就会一直阻塞着。
	*/
	var a chan int = make(chan int)
	//b := make(chan float32)
	go hellopara(a)
	fmt.Println("read channel before")
	c := <-a //读取信道的数值，因为是阻塞的，所以如果没有数值，会阻塞，就不需要sleep了
	//<-a 这样也是合法的
	fmt.Println("read channel after")
	fmt.Println(c)

	//信道死锁--只收不接（或只发不收）
	//<-a //提示deadlock

	//单向信道
	d := make(chan int)
	go chanOnlyRead(d) //双向信道可以作为参数赋值给单向信道，反过来就不行
	e := <-d
	fmt.Println(e)

	//关闭信道，用close，读取的时候可以接第二个返回值来判断该信道是否关闭
	f := make(chan int)
	go chanClose(f)
	g, ok1 := <-f //还能读出
	h, ok2 := <-f //读不出了
	fmt.Println(g, ok1, h, ok2)

	//利用循环自动接收信道数据，当没有数据时退出
	i := make(chan int)
	go chanForRange(i)
	for v := range i {
		fmt.Println(v)
	}

	//缓冲信道，只有信道数据满了，才会阻塞
	j := make(chan int, 2)
	j <- 1
	j <- 2 //写入时并没有阻塞
	//j <- 3 //再写会panic,必须等到消费才行
	fmt.Println(<-j, <-j)

	//信道的容量和长度,长度是指信道中当前排队的元素的个数
	fmt.Println(len(j), cap(j))

	//waitgroup 等待组，可以批量等待一批go协程，直到他们执行完毕
	//它通过引用计数来进行控制，引用计数减少为0才会继续执行
	var wg sync.WaitGroup
	wg.Add(1)
	go chanWaitGroup(&wg)
	wg.Add(1)
	//go chanWaitGroup(&wg)
	wg.Wait()//等待引用计数为0，否则会阻塞
	fmt.Println("WaitGroup end")
	

}

func hello() {
	fmt.Println("hello gotoutine!")
}

func hellopara(para chan int) {
	fmt.Println("write channel before")
	para <- 99 //写完后也会阻塞，直到有人读取数据
	fmt.Println("write channel after")
}

//singchan 唯送信道，只能往信道中扔数据
func chanOnlyRead(data chan<- int) {
	data <- 10
}

func chanClose(data chan int) {
	data <- 100
	close(data)
}

func chanForRange(data chan int) {
	data <- 101
	data <- 102
	close(data) //必须关闭，不然还会死锁
}

func chanWaitGroup(wg *sync.WaitGroup){
	time.Sleep(100 * time.Millisecond)
	wg.Done()//减少一次引用计数
}
