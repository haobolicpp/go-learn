package main

import "fmt"

/*
方法其实就是一个函数，在 func 这个关键字和方法名中间加入了一个特殊的【接收器】类型。
接收器可以是结构体类型或者是非结构体类型。接收器是可以在方法的内部访问的.

func (t Type) methodName(parameter list) {
}

自己理解：该功能可以模拟类，因为Go严格来说不是一门纯粹的面向对象语言

*/

//A struct
type A struct {
	key, value int
}

//值接收器 还是注意，参数不需要var
func (a A) setA(va int, vb int) {
	a.key = va
	a.value = vb
}

// 指针接收器
func (a *A) setAByPointer(va int, vb int) {
	a.key = va
	a.value = vb
}

func main() {
	a := A{1, 2}
	a.setA(3, 4) //传值
	fmt.Println(a)

	a.setAByPointer(100, 101) // 传指针
	fmt.Println(a)

	//结构体A的匿名字段--作为某方法接收器参数，A也可以直接调用该方法
	type B struct {
		value int
		A
	}
	b := B{1, A{8, 9}}
	fmt.Println("before b", b) //{1 {8 9}}
	b.setAByPointer(11, 12)
	fmt.Println("after b", b) // {1 {11 12}}

	//指针调用值接收器
	var d *A = &a
	d.setA(123, 456) //通过指针调用，其实go解释为了(*d).setA
	fmt.Println(d)

	//值调用指针接收器
	var e = A{90, 91}
	e.setAByPointer(0, 0) // go解释成了(&e).setAByPointer
	fmt.Println(e)

}
