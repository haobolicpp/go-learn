package main

import "fmt"

/*
在Go中，当一个类型定义了接口中的所有方法，我们称它实现了该接口。
*/

func main() {
	//接口的声明和实现
	a := Mystring("abcdefg")
	b := a.FindVowels()
	c := []rune("xyz") // tyoe 类型为int32
	fmt.Printf("a type:%T, b type:%T, c type:%T, valuea%s, valueb%c, valuec%c\n", a, b, c, a, b, c)

	//接口的实际用途举例,传入接口的切片（类似于基类指针）
	d := Dog{4}
	e := Cat{4}
	f := []Animal{d, e}
	fmt.Println("Dog and Cat legnums:", GetAnimalsLegs(f))
	fmt.Printf("dog type:%T\n", d)

	//接口对象，下面的赋值后，type变为了Dog,Value也是Dog的数值
	var g Animal
	g = d
	fmt.Printf("%T\n", g)

	//空接口，所有类型都实现了该接口 interface{}
	nilfunc(d)
	nilfunc(a)

	//类型断言，利用i.(T)实现获取接口底层的value
	h := 100
	assertfunc(h)
	//assertfunc("123")//运行报错：panic: interface conversion: interface {} is string, not int

	//类型选择 i.(type)
	findtype(h)
	findtype("123")
	findtype(d) //接口 type竟然是Animal，不是Myinterface，也就是说Dog实现了两个接口，但只体现了一个
	//体现的规则不清楚

	//实现接口：指针接收器和值接收器的不同
	var m Animal
	var n Chicken
	//m = n //错误,n只有指针接收器
	m = &n
	fmt.Println(m.GetLegsnum())

	//实现多个接口
	var o MyInterface
	var p Animal
	q := Dog{10} // Dog实现了多个接口的不同方法，可以赋值给不同的接口对象调用
	o = q
	p = q
	o.FindVowels()
	p.GetLegsnum()

	//接口的嵌套
	var mycombo ComboInterface
	mycombo = q
	mycombo.GetLegsnum()
	mycombo.FindVowels()

	//接口的nil
	var niltest Animal
	if niltest == nil {
		fmt.Println("niltest is nil")
	}
}

//MyInterface 定义一个接口
type MyInterface interface {
	FindVowels() []rune
}

//Mystring 定义自定义类型
type Mystring string

//FindVowels 实现该类型的接口MyInterface
func (str Mystring) FindVowels() []rune {
	var runeRet []rune
	for _, value := range str {
		if value == 'a' || value == 'e' || value == 'i' || value == 'o' || value == 'u' {
			runeRet = append(runeRet, value)
		}
	}
	return runeRet
}

//Animal 动物接口
type Animal interface {
	GetLegsnum() int //返回腿的数量
}

//Dog 狗
type Dog struct {
	legsnum int
}

//Cat 猫
type Cat struct {
	legsnum int
}

//Chicken 鸡
type Chicken struct {
	legsnum int
}

//FindVowels Dog可以实现多个接口
func (dog Dog) FindVowels() []rune {
	return []rune{1, 2, 3}
}

//GetLegsnum 狗实现获取腿接口
func (dog Dog) GetLegsnum() int {
	return dog.legsnum
}

//GetLegsnum 猫实现获取腿接口
func (cat Cat) GetLegsnum() int {
	return cat.legsnum
}

//GetLegsnum 指针接收器
func (chicken *Chicken) GetLegsnum() int {
	chicken.legsnum = 100
	return chicken.legsnum
}

//GetAnimalsLegs 动物腿求和
func GetAnimalsLegs(a []Animal) int {
	i := 0
	for _, v := range a {
		i += v.GetLegsnum()
	}
	return i
}

//nilfunc 空接口参数
func nilfunc(i interface{}) {
	fmt.Printf("i type:%T\n", i)
}

//类型断言接口 i.(T)
func assertfunc(i interface{}) {
	//不安全的方式
	v := i.(int)
	fmt.Println(v)

	//正确的方式
	v2, ok := i.(string)
	fmt.Println(v2, ok)
}

//findtype 类型选择
func findtype(i interface{}) {
	switch i.(type) {
	case string:
		fmt.Println("findtype string")
	case int:
		fmt.Println("findtype int")
	case Animal: //还可以和接口比较
		fmt.Println("findtype Animal")
	default:
		fmt.Println("findtype default")
	}
}

//ComboInterface 嵌套接口
type ComboInterface interface {
	Animal
	MyInterface
}
