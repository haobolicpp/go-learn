package main

/*
本部分包含了如下内容：
#数组、切片
*/

import "fmt"

func main() {
	TestArray()
	TestSlice()
}

//TestArray 数组测试，一定要有类型，且必须为同一个类型
func TestArray() {
	var a [3]int   //变量名单独的 后面是数组+类型
	fmt.Println(a) //结果[0 0 0]

	b := [3]int{1, 2, 3} //注意需要int
	fmt.Println(b)       //结果[1 2 3]

	c := [...]int{4, 5, 6} //省略个数
	fmt.Println(c)

	//var d [5]int
	//c = d //错误！数组的个数算是数组类型的一部分，不同类型不能赋值

	e := c //数值拷贝，修改e，c不会变化
	e[0] = 99
	fmt.Println("e[0]:", e[0], "c[0]:", c[0])

	//数组传函数参数-值传递
	testArrayFunc(b)
	fmt.Println("out func array[0] is:", b[0])

	//数组长度
	fmt.Println("array e len is:", len(e))

	//数组遍历（range的使用）
	fmt.Println("数组遍历(range):")
	for i, v := range e { //i是索引位置，v是该处的数值
		fmt.Println(i, v)
	}
	for _, v := range e { //省略下标，只获取数值
		fmt.Print(v)
	}
	fmt.Println("")

	//多维数组
	multiA := [2][2]string{
		{"abc", "efg"},
		{"1234", "xyzsss"},
	}
	for _, v := range multiA { // v代表第n行的数组
		fmt.Print(v, " ")
		for _, v2 := range v {
			fmt.Print(v2, " ")
		}
		fmt.Println("")
	}
	fmt.Println("传统方法：")
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			fmt.Println(multiA[i][j])
		}
	}
}

//本方式是按值传递的
func testArrayFunc(a [3]int) { //传递参数不需要var了
	a[0] = 100
	fmt.Println("in func array[0] is:", a[0])
}

//TestSlice 切片
//切片不存储元素，只是对原有数组的引用，对其修改都会反映到原始的数组上
func TestSlice() {
	//创建切片，直接var,后面用[]表示切片范围
	a := [5]int{1, 2, 3, 4, 5}
	var b = a[0:5] //下标应该从0开始,到5-1为止
	var c = a[5:5] //本值为空，相当于5~4
	fmt.Println("b:", b, "c:", c)

	//创建切片2
	d := []int{7, 8, 9}
	fmt.Println(d)

	//切片修改
	e := [...]int{1, 2, 3, 4, 5} //【注意】，此为数组
	fmt.Println("切片前e的数据:", e)
	var f = e[:]       // 创建一个包含所有元素的切片
	for i := range f { //---【注意】，i为下标，自动++，和c++不一样
		f[i]++
	}
	fmt.Println("切片修改后e的数据:", e)

	//多个切片可指向同一个数组
	g := e[:]
	h := e[:]
	g[0] = 100
	h[1] = 101
	fmt.Println("多切片修改后e的数据:", e)//100，101，3，。。。

	//切片的长度和容量
	i := e[1:3] //len2,cap4, 容量从1开始到最后
	j := i[:4]  //切片的切片，len4,cap4，容量从0开始，到3
	fmt.Println("i-len:", len(i), ",i-cap:", cap(i), ",j-len:", len(j), ",j-cap:", cap(j))

	//利用make创建切片
	k := make([]int, 10, 11) //参数一必须要有[]，参数二为长度，参数三为容量，参数三可省略且不能比参数二小
	fmt.Println(k)

	//追加切片元素，func append（s[]T，x ... T）[]T
	//append会创建一个新的切片，并将容量翻倍
	l := k[:]
	fmt.Println("l执行append前的cap:", cap(l)) //11
	l = append(l, 9, 10)
	fmt.Println("l执行append后的cap:", cap(l)) //24

	//零值切片
	var m []int
	if m == nil {
		m = append(m, 1, 2, 3)
		fmt.Println("零值nil，执行append后:", m)
	}

	//切片追加合并(利用...)
	n := []int{20, 21, 22}
	o := append(m, n...)
	fmt.Println("m、n切片合并后:", o)

	//切片的函数传递
	/*切片可看成如下形式：长度、容量和指向0位置元素的指针
	type slice struct {
		Length        int
		Capacity      int
		ZerothElement *byte
	}*/
	fmt.Println("n修改前:", n) //[20 21 22]
	sliceFunc(n)
	fmt.Println("n修改后:", n) //[101 21 22]

	//多维切片（类似多维数组）
	p := [2][3]int{
		{1, 2, 3},
		{4, 5, 6}, //注意逗号！！！！
	}
	fmt.Println("多维切片p:", p)

	//内存优化
	/*当一个切片指向一个数组时，本数组是不能回收的，可以调用copy(a,b)，将切片复制，然后
	原先指向的数组就可以被垃圾回收了
	*/
}

func sliceFunc(sliceA []int) {
	sliceA[0] = 101
}
