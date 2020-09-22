package main

import "fmt"

/*key-value*/

func main() {
	//初始化
	var amap map[string]int //默认nil，必须初始化才能用
	if amap == nil {
		amap = make(map[string]int)
	}

	var bmap = map[string]int{
		"123": 1,
		"456": 2,
	}
	fmt.Println(bmap)

	//添加-获取元素
	amap["baba"] = 250
	amap["mama"] = 666
	fmt.Println(amap, amap["mama"])

	//获取不存在的元素，返回对应元素的零值
	fmt.Println(amap["a"])

	//返回map中某个元素是否存在(双返回值，第二个返回值为true和false表示是否存在)
	value, ok := amap["baba"]
	fmt.Println(value, ok)

	//遍历
	for key, value := range amap {
		fmt.Println(key, value)
	}

	//删除某个元素--注意无返回值   delete(map, key)
	delete(amap, "baba")
	fmt.Println("删除元素baba后的：", amap)

	//遍历中删除
	cmap := map[int]int{
		1: 1,
		2: 2,
		3: 3,
		4: 4,
	}
	fmt.Println("cmap删除3前:", cmap)
	for key := range cmap {
		if key == 3 {
			delete(cmap, key)
		}
	}
	fmt.Println("cmap删除3后:", cmap)

	//map长度
	fmt.Println("cmap len:", len(cmap))

	//map是引用类型，和切片类似（类似qt-stl数据共享）
	dmap := cmap
	fmt.Println(dmap)

	//map的相等性
	//不能用==检查是否相等，只能检查是否是nil
	if dmap == nil {
		fmt.Println("dmap is nil")
	}
}
