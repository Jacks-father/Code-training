package main

import "fmt"

func main() {
	var a int = 1
	var b int = 2
	var c int
	//定义int数值变量
	c = (a * b)/2
	//公式
	fmt.Println("c的值是\t：\n",c)//打印c值
	fmt.Printf("c的数据类型是:%T",c)//打印c的数据类型
	//%T查看数据类型


	var d float32//定义float32浮点数值变量
	d = 3//赋值
	var e float32
	e = d * 3.14//圆周公式
	fmt.Println("e的圆周为：",e)
	fmt.Printf("周长e的数据类型是:%T\n",e)


	name := "刘述凯"//赋值string字符串类型
	age := 21//赋值int类型
	sex := "男"
	fmt.Println("我的名字是：\n",name)
	fmt.Println("我的年龄是：\n",age)
	fmt.Println("我的性别是：\n",sex)
	fmt.Printf("name的数据类型是：%T\n",name)
	fmt.Printf("age的数据类型是：%T\n",age)
	fmt.Printf("sex的数据类型是：%T\n",sex)
	//ctrl+z,复制上一行代码

	var f,g,h int
	f = 1
	g = 2
	h = 3
	fmt.Println(f,g,h)
	f = 4
	g = 5
	h = 6
	fmt.Println(f,g,h)

	var st1,st2,st3 = "刘",21,"凯"
	fmt.Println(st1,st2,st3)





}
