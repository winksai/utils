package main

import (
	"fmt"
	"net/http"
	"strconv"
	"sync"
	"utils/demo/global"
)

var WG sync.WaitGroup

func main() {

	go func() {
		for {
			select {
			case data1 := <-global.Ch1:
				if data1 == 1 {
					fmt.Println("执行了第一个")
				}
			case data2 := <-global.Ch2:
				if data2 == 1 {
					fmt.Println("执行了第二个")
				}
			case data3 := <-global.Close:
				if data3 {
					return
				}
			default:
			}
		}

	}()

	http.HandleFunc("/demo", Demo)
	http.HandleFunc("/demos", Demos)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
	/**
	1.CSP并发模型 协程和通道的运用
		等待组
		互斥锁  读写锁
		多路复用
		sync.map
	2.GMP调度模型
		当我们创建协程的时候，我们GO语言是如何进行调度的
	*/

	//1. 监听一个通道A，同时我还想监听一个通道B，如果A或者B有任意一个通道有值，都去执行

}

func Demo(_ http.ResponseWriter, r *http.Request) {
	data := r.FormValue("data")
	fmt.Println(data)
	dataInt, _ := strconv.Atoi(data)
	global.Ch1 <- dataInt
}

func Demos(_ http.ResponseWriter, r *http.Request) {
	data := r.FormValue("data")
	dataInt, _ := strconv.Atoi(data)
	global.Ch2 <- dataInt
}
