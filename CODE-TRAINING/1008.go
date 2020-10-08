package main

import "fmt"

func main() {
	var username string
	var password string
	fmt.Println("请输入用户名和密码：")
	fmt.Println("用户名：")
	fmt.Scanln(&username)
	fmt.Println("密码：")
	fmt.Scanln(&password)
	fmt.Printf("您输入的用户名是：%s,密码是：%s\n",username,password)//%s表示字符串

	if b :=7; b >= 1 {
		fmt.Printf("b的值为：%d\n",b)
		fmt.Println("b是一个正数")
	}//15到17行为b变量的作用范围，局限于15至17行的花框框，超出花框框打印就无效



	//switch语句
	var age int
	fmt.Println("请输入你的年龄：")
	fmt.Scanln(&age)
	switch age {
	case 18:
		fmt.Println("已满18，成人礼啦")
	case 20:
		fmt.Println("女生已达法定结婚年龄，男生未达")
	case 22:
		fmt.Println("男生女生都已达法定结婚年龄")
	default:
		fmt.Println("请等待……",age)
	}



	fmt.Println("请输入两个整数：")
	var num1 int
	var num2 int
	fmt.Scanln(&num1,&num2)

	fmt.Println("请输入要执行的操作符号（+，-，x,/）：")
	var str string
	fmt.Scanln(&str)

	switch str {
	case "+":
		fmt.Println("两数相加和为：",num1+num2)
		fallthrough  //fallthrough只能放在case分支的最后一行
	case "-":
		fmt.Println("两数相减差为：",num1-num2)
		break  //终止，结束语句向下执行
		fmt.Println("……")
	case "/":
		fmt.Println("两数相减差为：",num1/num2)
	default:
		fmt.Println("敬请期待……")
	}

	var num3 int
	var num4 int
	fmt.Println("请输入两个变量值：")
	fmt.Scanln(&num3,&num4)
	var str1 string
	fmt.Println("请输入要操作的符号（++，--，%）：")
	fmt.Scanln(&str1)
	switch str1 {//switch之后只能使用一个数值化语句，如：switch num := 5;num{……}
	case "++":
		num3++
		num4++
		fmt.Println("变量num3自增后的值为：",num3)
		fmt.Println("变量num4自增后的值为：",num4)
	}

	for i :=1;i <=20 ;i++{
		if i % 3 == 0 {
			fmt.Println(i)
		}
	}
}
