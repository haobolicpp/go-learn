package rectangle //属于什么包

import (
	"fmt"
	"math"
)

func init() {
	//在main函数执行前，本函数就调用了
	fmt.Println("restcalc init")
}

//Area 计算面积--首字母大写为导出
func Area(len, wid float64) float64 {
	return len * wid
}

//Diagonal 计算斜率
func Diagonal(len, wid float64) float64 {
	return math.Sqrt(len*len + wid*wid)
}

func interFunc() {

}
