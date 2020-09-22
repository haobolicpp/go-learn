package main

/*
error 是一个接口类型：
type error interface {
    Error() string
}
*/

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	//打开一个不存在的文件
	f, err := os.Open("/test.txt")
	if err != nil {
		fmt.Println(err) //打印 :open /test.txt: no such file or directory
	} else {
		fmt.Println(f.Name())
	}

	//获取更多的错误信息
	//Open返回的错误类型是*PathError
	// type PathError struct {
	// 	Op   string //操作Open等
	// 	Path string //错误文件的路径
	// 	Err  error //具体的错误信息
	// }
	//它通过声明Error()方法，实现error接口
	//pr := err.(*os.PathError) //基类转派生类
	if err, ok := err.(*os.PathError); ok { //if的用法
		fmt.Println(err.Path)
	}

	//当接口返回的错误类型有多种时，可以通过直接比较的办法
	files, err := filepath.Glob("[")
	if err != nil && err == filepath.ErrBadPattern {
		fmt.Println(err)
	} else {
		fmt.Println(files)
	}

	///自定义错误
	//方式一：使用New函数创建,该函数返回errorString对象
	a, err := myerrorByNew()
	fmt.Println(err, a)

	//方式二：使用Errorf格式化返回的错误信息
	err2 := myerrorByNewEF()
	fmt.Println(err2)

	//方法三：定义新的结构，重写接口error
	err3 := myerrorByerror()
	if e, ok := err3.(*myerror); ok { //基类转派生类，返回两个值
		fmt.Println(e.num)
	}

	//方法四：使用结构体类型的方法提供更多错误信息（其实是给方法三再添加新的方法）
	fmt.Println(err3.(*myerror).funcEx())

	//测试值接收器的方式实现error接口的Error方法
	err4 := myerror2Test()
	if e4, ok := err4.(myerror2); ok {
		fmt.Println(e4.leg)
	}
}

func myerrorByNew() (int, error) {
	return 1, errors.New("the error by New")
}

func myerrorByNewEF() error {
	return fmt.Errorf("%d,error", 1) //等同于errors.New(Sprintf(format, a...))
}

type myerror struct {
	errstr string
	num    int
}

//实现接口error中的Error()方法
func (e *myerror) Error() string {
	return fmt.Sprintf("重写接口error的Error方法，%d\n", e.num)
}

func myerrorByerror() error {
	return &myerror{"myerrorByerror", 99} // 指针接收器，需要取地址赋值
}

//给myerror实现更多的方法
func (e *myerror) funcEx() string {
	return "func EX"
}

//测试值接收器的方式
type myerror2 struct {
	err string
	leg int
}

//接口error实现方法Error
func (e myerror2) Error() string {
	return e.err
}

//返回myerror2的测试函数
func myerror2Test() error {
	errret := myerror2{"myerror2Test", 4}
	return errret
}
