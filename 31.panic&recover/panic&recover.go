package main

import (
	"fmt"
	"runtime/debug"
)

/*
需要注意的是，你应该尽可能地使用错误，而不是使用 panic 和 recover。
只有当程序不能继续运行的时候，才应该使用 panic 和 recover 机制
*/

func main() {
	/*panic的使用时机：
	1、发生了一个不能恢复的错误，此时程序不能继续运行。 一个例子就是 web 服务器无法绑定所要求的端口。
	在这种情况下，就应该使用 panic，因为如果不能绑定端口，啥也做不了。
	2、发生了一个编程上的错误，比如传入的参数非法。
	panic内建签名：当程序终止时，会打印传入 panic 的参数
	func panic(interface{})
	*/
	//testpanic(nil) //程序会退出，并打印堆栈信息

	//有延迟函数的panic,会执行完延迟函数后退出（协程也是如此--未测试）
	//testpanicdefer(nil)

	/*recover*/
	/*
		recover 是一个内建函数，用于重新获得 panic 协程的控制.
		func recover() interface{}
		只有在延迟函数的内部，调用 recover 才有用。在延迟函数内调用 recover，可以取到 panic 的错误信息，并且停止 panic 续发事件（Panicking Sequence），程序运行恢复正常。
		如果在延迟函数的外部调用 recover，就不能停止 panic 续发事件。
		只有在相同的 Go 协程中调用 recover 才管用。
	*/
	a := testrecover(nil)
	fmt.Println(a) //返回初始值0

	//运行时panic及捕获
	//主要针对运行时错误，如数组越界，实际是调用了内置的panic
	testruntimepanic() //打印：recover a mypanic: runtime error: index out of range

	//恢复后获得堆栈跟踪信息 --利用debug.PrintStack
	testPrintStack()
}

//测试panic,传入的参数非法
func testpanic(i *int) {
	if i == nil {
		panic("testpanic input i is nill")
	}
}

//延迟函数panic
func testpanicdefer(i *int) {
	defer fmt.Println("testpanicdefer")
	if i == nil {
		panic("testpanic input i is nill")
	}
}

//recover--
func deferrecover() {
	if r := recover(); r != nil {
		fmt.Println("recover a mypanic:", r)
	}
}

//recover test
func testrecover(i *int) int {
	defer deferrecover()
	if i == nil {
		panic("testrecover panic")
	}
	return 1
}

// runtime panic 及recover
func testruntimepanic() {
	defer deferrecover()
	a := []int{1, 2}
	a[3] = 0
}

// PrintStack
func deferPrintStack() {
	if a := recover(); a != nil {
		fmt.Println("recover deferPrintStack:", a)
		debug.PrintStack()
	}
}
func testPrintStack() {
	defer deferPrintStack()
	panic("testPrintStack panic")
}
