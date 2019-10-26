package main

import (
	"sync"
	"time"
)

var myRes = make(map[int]int, 100)

var lock sync.RWMutex

func main() {

	for i := 1; i < 100; i++ {
		go last(i)
	}

	time.Sleep(1 * time.Second)

	for i := 1; i < 100; i++ {
		println(myRes[i])
	}
}

func last(n int) {
	res := jiecheng(n)
	lock.Lock()
	intoMap(n, res)
	lock.Unlock()
}


func jiecheng(n int) int {
	var res = 1

	for i := 0; i < n; i++ {
		res *= i + 1
	}

	return res

}

func intoMap(n int, res int) {
	myRes[n] = res
}