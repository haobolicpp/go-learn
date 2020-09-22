package main

import "fmt"

/*
含有 defer 语句的函数，会在该函数将要返回之前，调用另一个函数.
相当于延迟调用，可以是函数或方法。
*/

func main() {
	//函数调用
	test()

	//defer的参数赋值时机，并不是在最后赋值的，而是根据defer的位置
	testpara()

	//defer栈，多次调用defer时，会按照后入先出的顺序调用defer
	testdeferstack()

}

func defertest() {
	fmt.Println("defer func run")
}

func test() {
	defer defertest()
	fmt.Println("test run...")
}

func deferpara(i int) {
	fmt.Printf("defer %d\n", i)
}

func testpara() {
	i := 100
	fmt.Printf("defer before:%d\n", i)
	defer deferpara(i) //打印100

	i = 99
	fmt.Printf("func end:%d\n", i)
}

func testdeferstack() {
	defer deferpara(999)
	defer defertest()
}
