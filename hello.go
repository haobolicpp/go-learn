// 该文件包含如下基础内容
//- 输出
//- 编译go、执行go
//- 变量定义 var :=
//- 常量定义
//- 函数
//- 自定义包及导入
//- if else for switch关键字
//- os.Args[] 0是全路径 后面是输入的内容，比如go run hello.go xxx，那么os.Args[1]就是xxx
package main

import (
	"fmt"
	"golearn/rectangle"
	"math"
	"os"
	"unsafe"
)

var _ = math.Sqrt //错误屏蔽器，虽然没有调用math，但可通过这种方式导入math包

//普通函数
func testDemo(price, no int) int {
	return price * no
}

//多返回值函数
func testMultiRet(para int) (int, float32) {
	para = 1
	return para, float32(para) / 10.0
}

//命名返回值
func testNameRet() (ret1, ret2 int) {
	ret1 = 100
	ret2 = ret1 * 100
	return
}

func main() {
	fmt.Println("hello world")

	/*直接执行go文件*/
	//go run hello.go

	/*编译go文件*/
	//go build hello.go

	/*变量*/
	var age int //int型，默认为0
	fmt.Println("my age default", age)
	age = 29
	fmt.Println("my new age", age)

	var height int = 178 //警告是可以自动推断
	var height1 = 180
	fmt.Println("my height is", height, height1)

	//var a, b = 1, 2 //多个变量
	//声明不同类型的变量
	var (
		name   = "bob"
		bobage = 1
		h      int
	)
	fmt.Println("my name is", name, "bob age is", bobage, h)
	//简短声明（没有var、类型）:= go 会自动推断
	names, anges := "bob_s", 12
	a, e := 1, 2
	//a, e := 3, 4
	a, c := 3, 4 //左边至少一个未声明过
	//a = "123"//error go是强类型语言，不允许某一类型的变量赋值为其他类型的值
	fmt.Println(a, e, c, names, anges)

	/* 类型
	bool
	int8, int16, int32, int64, int //有符号整型,最后的int在32位下是32位，64位下是64位
	uint8, uint16, uint32, uint64, uint
	float32, float64
	complex64, complex128
	byte 是 uint8 的别名。
	rune 是 int32 的别名
	string
	*/
	var boola = true
	boolb := false
	fmt.Println(boola, boolb)

	var ia int = 1
	var ib int8 = 10
	fmt.Printf("a type:%T, b type:%T, a size:%d,b size：%d哈哈\n", ia, ib, unsafe.Sizeof(ia), unsafe.Sizeof(ib)) //%T类型

	//不同类型变量相加
	ifirst := 1
	fsec := 2.0
	comp := ifirst + int(fsec) //用T(x)强转
	fmt.Println(comp)

	//变量赋值
	type mystring string
	var sfuzhi = "123"
	//var mys mystring = sfuzhi //重新定义的类型也不能强转
	var mys mystring = mystring(sfuzhi)
	fmt.Println(mys)

	/*常量*/
	const ca = 1 //不需要类型
	//const cb = math.Abs(-1)
	var hello = "hello world"
	const chello string = "c"
	const chellos = "cs"
	fmt.Printf("%T,%T,%T\n", ca, hello, chello)

	//常量赋值特例
	var intVar int16 = ca
	var float32Var float32 = ca
	fmt.Println(intVar, float32Var)

	/*函数*/
	fmt.Println("testDemo ret: ", testDemo(1, 2)) // 普通函数调用
	para, floatv := testMultiRet(4)
	fmt.Println("testMultiRet :", para, floatv) // 多返回值函数调用
	ret1, ret2 := testNameRet()
	fmt.Println("testNameRet:", ret1, ret2) // 命名返回值，里面直接return
	ret3, _ := testMultiRet(1)              //_ 表示空白符，可以表示任意类型
	fmt.Println("aaa", ret3)

	/*自定义包*/
	area := rectangle.Area(3, 4)
	dia := rectangle.Diagonal(3, 4)
	fmt.Println(area, dia)

	/*if-else  for语句*/
	if ca <= 1 {
		fmt.Printf("")
	} else if ca > 1 { //注意else必须在}后面
		fmt.Printf("")
	}

	for i := 1; i < 10; i++ {
		if i == 9 {
			break
		} else {
			continue
		}
	}
	j := 1
	for {
		j++
		if j > 5 {
			break
		}
	}

	/*switch语句*/
	switchnum := 4
	switch switchnum {
	case 1:
		fmt.Println("case 1") // 默认执行完成后跳出
	case 2, 3, 4:
		fmt.Println("case 2,3,4") // 多判断语句
		fallthrough               //不跳出，继续下面的
	case 5:
		fmt.Println("case 5")
	default:
		fmt.Print("default")
	}

	/*输入*/
	var strInput string
	fmt.Scanln(&strInput)
	fmt.Println(strInput)

	/*从main接收参数*/
	fmt.Println(os.Args[0]) //Args[0]打印的是exe的全路径
	if len(os.Args) > 1 {
		arg := os.Args[1]
		fmt.Println(arg)
	}

}
