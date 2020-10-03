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


	//变量的集合
	var (
		name = "liushukai"
		age = 20
		sex = "男"
		address = "江西省九江市那嘎达的"
	)

	fmt.Println(name,age,sex,address)

	//舍弃变量
	num01,_ := 9,7
	fmt.Println(num01)

}
