package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//for 循环
	res :=1
	for i :=1;i <=5 ; i++ {
		res *=i
	}
	fmt.Println("5的阶乘：",res)

	var x int
	for ;x < 12 ;  {//for 表达式1；表达式2；表达式3{……}，表达式1缺失，初始化时什么都不做，允许；表达式2缺失，默认为true，表示满足条件，会执行循环体；表达式3缺失，什么都不操作，允许
		fmt.Println(x)//若表达式2都缺失，则无限次打印内容
		x++
	}




	//多层循环
	for num1 :=0; num1 < 5 ; num1++ {//外行控制行数
		for a :=0;a < 5 ; a++ {//内行控制列数
			fmt.Print("hello\t")
		}
		fmt.Println()
	}

	//99乘法表
	for z :=1; z <=9 ; z++ {
		for y :=1; y <= z ; y++ {
			fmt.Printf("%d * %d = %d\t",z,y,z*y)
		}
		fmt.Println()
	}


	for a :=1; a <= 5 ; a++ {
		for b :=1; b <= 5-a; b++ {
			fmt.Print(" ")//print不会换行
		}
		for star := 1; star <=2*a-1; star++ {
			fmt.Print("*")
		}
		fmt.Println()
	}


	//打印椎体
	for a :=1; a <=4; a++ {
		for b :=1; b <=a; b++ {
			fmt.Print(" ")
		}
		for star :=1; star <= 9 - 2 * a; star++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
	
	
	
	count :=0  //记录打印的次数
	for f :=0 ; f <=100; f++ {
		if f % 6 ==0 && f % 4 ==0 {
			fmt.Println(f)
			count++  //每次打印，就代表找到一个目标值
			if count == 6 {
				break
			}
		}
	}


	for g :=0 ; g <=100; g++ {
		if g <=30 || g >= 90 {
			continue
		}
		if g % 6 ==0 || g % 4 ==0 {
			fmt.Println(g)
			}
		}

	OUT:for j :=1; j <=10; j++ {
		for k :=1; k <=6; k++ {
			if k ==3{
				break OUT
			}
			fmt.Println("j",j,"k",k)
		}
	}
	fmt.Println("over")



	var num001 int
	for num001 :=0; num001 < 1000; num001++ {
		if num001 < 10 {
			if num001 * num001 *num001 == num001{
				fmt.Println(num001)
			}
		}
		if num001 >=100 && num001 <=999{
			if num001 * num001 *num001 ==num001 {
				fmt.Println(num001)
			}
		}
	}
	fmt.Println(num001)


	//空心三角
	for l :=1;l <= 6; l++{
		for o :=1; o <=l; o++ {
			if l ==6 {
				fmt.Print("* ")
				continue
			}
			if o ==1 || o ==l {
				fmt.Print("* ")
			}else {
				fmt.Print("  ")
			}
		}
		fmt.Println()
	}



	//goto 语法：强制跳转到goto语句所指定行进行程序的执行
	var p = 20
	for p < 30 {
		if p ==25 {
			p +=1
			goto TOP1
		}
		fmt.Printf("p的值：%d\n",p)
		p++
	}
	TOP1:
		fmt.Println("OVERGAME")



	//随机数(rand.Intn(n:5):int表示生成的随机数是一个整数，5表示生成随机数的范围：0.1.2.3.4)
	n := rand.Intn(5)//intn（n int）获取随机整数，区间范围为前闭后开
	fmt.Println(n)


	//时间戳：获取当前时间和时间戳
	time1 := time.Now()
	//time.Unix()time1时间距离1970年1月1日，格林尼治时间
	//time.UnixNano()  //距离1970年1月1日格林尼治标准时间的参数
	rand.Seed(time1.UnixNano())//默认种子值1



	//
	var n1 int
	TOP2:
		fmt.Println("请输入一个整数数字：")
		fmt.Scanln(&n1)
	rand.Seed(time.Now().UnixNano())
	n2 := rand.Intn(n1*2)
	if n1 > n2 {
		fmt.Println("输入的数字比生成数大，程序结束")
	}else {
		fmt.Println("不满足条件，程序执行")
		goto TOP2
	}
}
