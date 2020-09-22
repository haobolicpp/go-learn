package main

/*可变参数函数*/

import "fmt"

func main() {
	//后面的1, 2, 3, 4, 5, 6参数被编译为切片传入
	a, b := findnumindex(5, 1, 2, 3, 4, 5, 6)
	fmt.Println(a, b)

	//给可变参数函数传入切片，用...
	c := []int{5, 6, 7, 8, 9}
	d, e := findnumindex(6, c...)
	fmt.Println(d, e)

	useslice(c...)
	fmt.Println(c)

}

//第二个参数为可变参数
//参数名...类型（只能描述一种类型）
func findnumindex(num int, nums ...int) (int, int) {
	for i := range nums {
		if num == nums[i] {
			return i, 0
		}
	}
	return -1, -1
}

func useslice(nums ...int) {
	nums[0] = -1
	nums = append(nums, 100) //append会创建一个新的切片，且容量翻倍
	fmt.Println(nums)
}
