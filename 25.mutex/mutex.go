package main

import (
	"fmt"
	"sync"
)

/*
通过 Mutex 和信道来处理多协程访问静态资源的竞态条件
*/

var x int
var mtx sync.Mutex

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go increaceNum(&wg)
	go increaceNum(&wg)
	wg.Wait()
	fmt.Println(x)
}

func increaceNum(wg *sync.WaitGroup) {
	mtx.Lock() //进行同步
	x++
	mtx.Unlock()
	wg.Done()
}
