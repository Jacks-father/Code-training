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
	fmt.Println("程序结束",age)


	num := 17
	fmt.Scanln(&num)
	fmt.Println(num)
	if num % 2 ==0 {
		fmt.Println("偶数")
	}else {
		fmt.Println("对不起，未被除尽，奇数")
	}
	fmt.Println("程序结束",num)







}
