package main

import (
	"fmt"
)

func main() {
	//关系运算符
	//关系运算符的运算结果是布尔类型
	//<,>,==,<=,>=,!= 等符号
	sum1 :=10
	sum2 :=20
	result1 :=sum1 < sum2//比较大小
	result2 :=sum1 > sum2
	fmt.Println(result1,result2)//输出结果是大是小
	sum3 :=10
	result3 :=sum1 == sum3
	result4 :=sum1 >= sum3
	result5 :=sum1 <= sum3
	result6 :=sum1 != sum3
	fmt.Println(result3,result4,result5,result6)



	//逻辑运算符
	//基本操作符：“与”操作：&&;规律是：全真则真，一假则假
	//			 “或”操作：||;规律是：一真则真，全假则假
	//			 “非”操作：！
	a1 := true
	a2 := false
	result7 :=a2 && a1//“与”操作&&
	fmt.Println(result7)//规律：全真则真，一假则假

	a3 := true
	a4 := false
	result8 :=a4 || a3//“或”操作||
	fmt.Println(result8)//规律：一真则真，全假则假

	a5 := 6
	a6 := 5
	result9 :=!(a6 < a5)//“非”操作！
	fmt.Println(result9)//规律： 与比较结果相反

	//综合题
	a :=1
	b :=2
	c :=3
	result :=!(a < b && b/a > a && c > b || c != 5)
	fmt.Println("result的运算结果是：\n",result)


	num01 :=5	//5-->101	//十进制转换为二进制，位数不够在开头补0
	num02 :=6	//6-->110
	res :=num01 & num02		//二进制位比较：101
							//             110
							//res：		   100    // 按位与&运算：0和1比较输出0,1和1比较输出1

	res01 :=num01 | num02	//	101
							//	110
							//	111			// 按位或|运算：0和1比较输出1,1和1比较输出1

	res02 :=num01 ^ num02	//	101
							//	110
							//	011			//按位异或^运算：数字比较，相同输出0，不同输出1

	fmt.Println(res,res01,res02)
}
