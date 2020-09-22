package main

import (
	"fmt"
	"unicode/utf8"
)

//Go 中的字符串是兼容 Unicode 编码的，并且使用 UTF-8 进行编码

func main() {
	//示例
	astr := "hello string"
	fmt.Println(astr)

	//获取字符串的每一个字节 --十六进制
	for i := 0; i < len(astr); i++ {
		fmt.Printf("%x ", astr[i])
	}
	fmt.Println()

	//打印字符
	bstr := "爱你祖国"
	for i := 0; i < len(bstr); i++ {
		fmt.Printf("%c ", bstr[i]) //输出乱码，默认UTF-8编码，一个汉字可能由多个字节
	}
	fmt.Println("len(bstr):", len(bstr))

	//利用rune打印UTF-8(rune 是 Go 语言的内建类型，它也是 int32 的别称
	runes := []rune(bstr) //注意是括号
	for i := 0; i < len(runes); i++ {
		fmt.Printf("%c ", runes[i])
	}
	fmt.Println("len(runes):", len(runes))

	//for range循环
	cstr := "Señor"
	for index, value := range cstr {
		//结果下标是0，1，2，4，5.
		fmt.Printf("for-range,index:%d, value:%c\n", index, value)
	}

	//字节切片构造string
	byteslice := []byte{0x43, 0x61, 0x66, 0xC3, 0xA9}
	dstr := string(byteslice)
	fmt.Println(dstr)

	//字符串长度
	fmt.Printf("len:%d, runecount:%d\n", len(cstr), utf8.RuneCountInString(cstr))

	//字符串修改（切片修改）
	estr := "123"
	//estr[0] = 'a' //无法修改
	erune := []rune(estr)
	erune[0] = 'a'
	fstr := erune
	fstr[1] = 'b' //修改的是同一个数据
	fmt.Println(estr, erune, fstr)
}
