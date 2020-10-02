package main

import "fmt"

func main() {
	num1,num2 := 3,5//赋值
	fmt.Println(num1 + num2)//打印数值和

	num1,num2 = 6,10
	fmt.Println(num1,num2)

	num1,num3 := 5,8
	fmt.Println(num1,num3)

	num1 = 10//赋值
	fmt.Println(num1)
}
