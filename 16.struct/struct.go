package main

import "fmt"

//Employee 职员
type Employee struct {
	firstname, lastname string
	age, salary         int
}

func main() {
	//创建命名的结构体
	a := Employee{
		firstname: "li",
		age:       29,
		salary:    10000,
		lastname:  "bob",
	}

	b := Employee{"li", "bob", 29, 10000}

	fmt.Println(a, b)

	//创建匿名结构体--直接用struct
	c := struct {
		firstname, lastname string
		age, salary         int
	}{
		"li", "bob", 29, 10000,
	}
	fmt.Println(c)

	//零值初始化1
	var d Employee
	//零值初始化2
	e := Employee{
		firstname: "li",
	}
	fmt.Println(d, e)

	//结构体指针及成员访问
	f := &Employee{"li", "bb", 1, 2}
	fmt.Println(f.firstname, (*f).salary)

	//匿名字段，结构体字段只有类型
	type person struct {
		string
		int
	}

	g := person{"tom", 50}
	fmt.Println(g, g.string, g.int) // 可以通过类型访问

	//提升字段（Promoted Fields）(内嵌匿名结构体)
	type A struct {
		key, value int
	}
	type B struct {
		value2 int
		A
	}
	h := B{1, A{2, 3}}
	fmt.Println(h.key) //可以直接访问

	//结构体导出
	//通过定义结构体名称首字母大写，来导出本包的结构体

	//结构体是否相同比较
	i := person{"a", 1}
	j := person{"a", 1}
	if i == j {
		fmt.Println("i和j相同")
	}

	//里面如果包含map，则不能比较是否相同
}
