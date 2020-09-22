package main

import "fmt"

/*利用 *T 来指向一个T类型的变量*/

func main() {
	//基本
	a := 1
	var pb *int = &a
	pc := &a
	fmt.Println(pb, a, pc)

	//解引用
	d := *pc
	fmt.Println(d)

	//向函数传递指针
	e := 99
	testpoint(&e)
	fmt.Println(e) //被修改为了100

	//go不推荐通过指针向一个函数传递 数组,这里省略了传指针的做法
	f := [3]int{1, 2, 3}
	testpara(f[:])
	fmt.Println(f)

	//go 不支持指针的运算
	//pc++
}

func testpoint(val *int) {
	*val = 100
}

func testpara(val []int) {
	val[0] = 100
}
