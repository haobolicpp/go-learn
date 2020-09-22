package main

import (
	"fmt"
	"reflect"
)

/*
@@什么是反射？
    在 Go 语言中，reflect 实现了运行时反射。reflect 包会帮助识别 interface{} 变量的底层【具体类型和具体值】。
@@为何需要检查变量，确定变量的类型？
@@reflect 包
   reflect.Type 和 reflect.Value
    reflect.Kind
@@NumField() 和 Field() 方法
@@Int() 和 String() 方法
@@完整的程序
@@我们应该使用反射吗？
*/

type tMystruct struct {
	a int
	b string
}

func main() {
	//reflect.Type 和 reflect.Value
	//reflect.Typeof 表示 interface{} 的具体类型，而 reflect.Valueof 表示它的具体值
	a := tMystruct{1, "2"}
	fmt.Println("reflect.TypeOf:", reflect.TypeOf(a).String()) //main.tMystruct
	fmt.Println("reflect.Valueof:", reflect.ValueOf(a))        //{1, "2"}

	//reflect.TypeOf().Kind : 返回数据类型如struct/int
	fmt.Println("reflect.Kind:", reflect.TypeOf(a).Kind()) //struct
	b := 1
	fmt.Println("reflect.Kind:", reflect.TypeOf(b).Kind()) //int

	//NumField() 和 Field() 方法
	//reflect.Valueof().NumField()  方法返回结构体中字段的数量
	//reflect.Valueof().Field(i int) 方法返回字段 i 的 reflect.Value。
	fmt.Println("reflect.Valueof().NumField():", reflect.ValueOf(a).NumField())
	fmt.Println("reflect.Valueof().Field():", reflect.ValueOf(a).Field(0))
	fmt.Println("reflect.Valueof().Field():", reflect.ValueOf(a).Field(1))
}
