package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//判断随机数
	for {
		var n1 int
		fmt.Println("请输入一个整数数字：")
		fmt.Scanln(&n1)
		rand.Seed(time.Now().UnixNano())
		n2 := rand.Intn(n1 * 2)
		if n1 > n2 {
			fmt.Println("输入的数字比生成数大，程序结束")
			break
		} else {
			fmt.Println("不满足条件，程序执行")
		}
	}





}
