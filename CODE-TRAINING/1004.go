package main

import "fmt"

func main() {
		a := 3
		b := 5
		c := 7
		d := 9
		sum := a + b
		sum1 := a - b
		sum2 := c * d
		sum3 := c / d//算术运算符
		fmt.Println(sum,sum1,sum2,sum3)
		sum++//自增
		sum1++
		sum2--//自减
		sum3--
		fmt.Println(sum)
		fmt.Println(sum1)
		fmt.Println(sum2)
		fmt.Println(sum3)
		sum4 :=sum + sum1 + sum2 + sum3
		fmt.Println(sum4)


		class0 := 8
		class1 := 10
		result := class0 > class1
		fmt.Println(result)




}
