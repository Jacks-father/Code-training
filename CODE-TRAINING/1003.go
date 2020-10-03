package main

import "fmt"

var num  = 3//全局变量

func main() {

	var num  = 10
	fmt.Println(num)


	const PAI float32  = 3.1415926//语法：const 常量名字 数据类型 = 内容值
	//被定义的常量名用大写字母
	fmt.Println(PAI)
	//PAI = 3.1415，被const表明的常数不可被更改


	const BAIDU  = "http://www.baidu.com"//常量名网址
	fmt.Println(BAIDU)

	const (
		age = 100
		sex = "人妖"
		Monday = "星期一"
		Tuesday = "星期二"
		Thursday = "星期三"
	)
	fmt.Println(age,sex,Monday,Tuesday,Thursday)



	var num01 int8
	num01 = 001
	var num02 int64
	num02 = 007
	//sum := num01 + int8(num02)
	sum := int64(num01) + num02//将不同int类型转换统一int类型求和
	fmt.Println(sum)


	grade := 99.999
	result := int64(grade)//转换类型
	fmt.Println(result)
	fmt.Printf("result的数据类型是:%T\n",result)

	var x uint8 = 6//uint是byte的别名
	var y byte = 8
	sum01 := x + y
	fmt.Println(sum01)

	var PAI01 = 3.141592654
	fmt.Printf("PAI01的数值是：%f\n",PAI01)
	fmt.Printf("保留PAI01小数点后三位的数值是：%.3f\n",PAI01)//保留小数点后几位用：. + 要保留几位数 + f



	num00 := "saddadaasdad"
	name := "小明不叫晓明"//一个汉字相当于三个字母
	length := len(name)
	fmt.Println("num00的长度是：",len(num00))
	fmt.Println("字符串name的长度是：",length)
}
func hello()  {
	fmt.Println(num)
}