package main

import "fmt"

func main() {
	//赋值运算
	var a int
	a = 1
	fmt.Println(a)
	a += 3
	fmt.Println(a)
	a *= 3
	fmt.Println(a)
	a /= 3
	fmt.Println(a)
	a &= 3
	fmt.Println(a)
	a ^= 3
	fmt.Println(a)
	a |= 3
	fmt.Println(a)
	a %= 3
	fmt.Println(a)
	a >>= 3
	fmt.Println(a)
	a <<= 3
	fmt.Println(a)


	num1 :=2
	num2 :=3
	fmt.Println(num1,num2)
	num3 := num2
	num1 = num2
	num2 = num3
	fmt.Println(num1,num2,num3)

	num4,num5 := 6,8
	fmt.Println(num4,num5)
	num4,num5 = num5,num4
	fmt.Println(num4,num5)


	//流程语句
	var grade int
	grade = 60
	if grade>=60 {
		fmt.Printf("成绩为：%d，成绩合格\n",grade)
	}
	fmt.Println("程序结束\n",grade)

	//if...else语句
	age := 17
	if age>=18 {
		fmt.Println("恭喜你成为了一名大人\n",age)
	}else {
		fmt.Println("对不起，你还未满18，请勿出入儿童不宜的场所 \n",age)
	}
	fmt.Println("程序结束")


	num := 17
	fmt.Scanln(&num)//可在运行程序中输入内容
	fmt.Println(num)
	if num % 2 ==0 {
		fmt.Println("偶数")//符合num % 2 == 0，则输出该内容
	}else {
		fmt.Println("对不起，未被除尽，奇数")//不符合，则输出该内容
	}
	fmt.Println("程序结束")

	var a1 int
	var b1 int
	fmt.Println("请分别输入a1和b1的值：")
	fmt.Scanln(&a1,&b1)//可在运行程序中输入内容
	fmt.Printf("a1变量的值为：%d，b1变量的值为：%d\n",a1,b1)
	fmt.Println("a1的内存地址是：",&a1)//通过&符号查看内存地址
	fmt.Println("b1的内存地址是：",&b1)
	fmt.Scanf("%d,%d",&a1,&b1)//输出，注意%d与%d之间的逗号，键盘输入时也要用逗号隔开
	fmt.Printf("a1变量的值为：%d，b1变量的值为：%d\n",a1,b1)
	fmt.Println("a1的内存地址是：",&a1)//通过&符号查看内存地址
	fmt.Println("b1的内存地址是：",&b1)

	var age01 int
	fmt.Println("请输入你的年龄：")
	fmt.Scanln(&age01)
	var sex int
	fmt.Println("请输入你的性别：")
	fmt.Scanln(&sex)
	if sex ==1 {//符合 sex == 1，则进入下列程序
		if age01 >=22 {
			fmt.Println("可以结婚了，但先得有个对象哦")//符合年龄，则打印该内容
		}else {
			fmt.Println("拜拜了你勒，孩纸还小，别想太多哦")//不符合，则打印该内容
		}
	}else {//不符合 sex == 1，则执行下列程序
		if age01 >=20 {
			fmt.Println("找个新东方厨师，就嫁了吧")//符合年龄，则打印该内容
		}else {
			fmt.Println("小妹妹还需发育哦")//不符合，则打印该内容
		}
	}



	//if 语句扩展
	var sex0 int
	fmt.Println("请输入性别：")
	fmt.Println("友情提示：1表示男，2表示女，3表示人妖，4表示未知性别，请随意填写自己的性别，后果自负哦")
	fmt.Scanln(&sex0)
	if sex0 == 1 {
		fmt.Println("性别为男")
	}else if sex0 == 2 {
		fmt.Println("性别为女")
	}else if sex0 == 3 {
		fmt.Println("性别为人妖")
	}else if sex0 == 4 {
		fmt.Println("性别为未知")
	}





}
