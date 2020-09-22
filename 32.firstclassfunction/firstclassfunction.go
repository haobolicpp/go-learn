package main

import "fmt"

/*--头等函数
@@支持头等函数（First Class Function）的编程语言，
可以把函数赋值给变量，也可以把函数作为其它函数的参数或者返回值。Go 语言支持头等函数的机制。
@@【匿名函数】
@@【用户自定义的函数类型】
@@【高阶函数】
@@【在其它函数中返回函数】
@@【闭包】
@@【实际用途】
*/
func main() {
	//@@【匿名函数】
	a := func() {
		fmt.Println("hello world first class function")
	}
	a()
	fmt.Printf("a tyope:%T\n", a)

	//@@【匿名函数】-不赋值变量
	func() {
		fmt.Println("【匿名函数】-不赋值变量")
	}()

	//@@【匿名函数】-传递参数
	func(s string) {
		fmt.Println("【匿名函数】-传递参数:", s)
	}("hello")

	//@@【用户自定义的函数类型】
	var b add = func(c int, d int) int {
		return c + d
	}
	fmt.Println("【用户自定义的函数类型】-add:", b(1, 2))

	//@@【高阶函数】
	/*	1/接收一个或多个函数作为参数
		2/返回值是一个函数*/
	//@@【高阶函数】-函数作为参数
	c := func(i int, j int) int {
		return i + j
	}
	funcpara(c)

	//@@【高阶函数】-函数作为返回值
	d := funcret()
	fmt.Println(d(3, 4))

	//@@【闭包】
	//闭包（Closure）是匿名函数的一个特例。当一个匿名函数所访问的变量定义在函数体的外部时，就称这样的匿名函数为闭包。
	e := 100
	f := func() {
		fmt.Println(e) //访问外部变量
	}
	f()

}

//自定义函数类型签名--add
type add func(a int, b int) int

//函数作为参数
func funcpara(b func(a int, b int) int) {
	fmt.Println(b(1, 2))
}

//函数作为返回值
func funcret() func(a int, b int) int {
	f := func(a, b int) int {
		return a + b
	}
	return f
}
