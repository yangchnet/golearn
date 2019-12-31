/**
 * @Description: go语言圣经第二章
 * @author lichang
 * @date 19-12-25
 * @time 下午3:07
 */
package main

import (
	"chap2/lengthconv"
	"chap2/popcount"
	"chap2/tempconv"
	"fmt"
)

const boilingF = 212.0

func main() {
	// const x = 14
	// const n = 5
	// fmt.Printf("%d\n",fibSum(n))
	fmt.Printf("Brrr! %v\n", tempconv.AbsoluteZeroC)
	fmt.Printf("LichanHigh: %v\n", lengthconv.LichangHigh)
	a := popcount.PopCount(5)
	fmt.Printf("%v\n", a)
}

func boiling() {
	/**
	 * @Auther:lichang
	 * @Date: 19-12-25 下午3:18
	 * @Description: 华氏度转为摄氏度
	 * @Param:
	 * @Return:
	 */
	var f = boilingF
	var c = (f - 32) * 5 / 9
	fmt.Printf("boiling point = %g or %g\n", f, c)
}

/*****************Dividing line*****************/

func fToC(f float64) float64 {
	/**
	 * @Auther:lichang
	 * @Date: 19-12-25 下午3:18
	 * @Description: 将华氏度转为摄氏度
	 * @Param: 华氏度
	 * @Return: 摄氏度
	 */
	return (f - 32) * 5 / 9
}

/*****************Dividing line*****************/

func incr(p *int) int {
	/**
	 * @Auther:lichang
	 * @Date: 19-12-25 下午3:32
	 * @Description: C语言i++的模拟，但不推荐
	 * @Param: p
	 * @Return: p+1
	 */
	*p++
	return *p
}

/*****************Dividing line*****************/

func gcd(x, y int) int {
	/**
	 * @Auther:lichang
	 * @Date: 19-12-25 下午4:32
	 * @Description: 求两个数的最大公约数 （欧几里得平凡数算法）
	 * @Param: 两个整数
	 * @Return: 最大公约数
	 */

	//for y !=0 { 标准写法
	//	y = x % y
	//	x = y
	//}
	for y != 0 { //简单写法
		x, y = y, x%y
	}
	return x
}

/*****************Dividing line*****************/

func fib(n int) int {
	/**
	 * @Auther:lichang
	 * @Date: 19-12-25 下午4:54
	 * @Description: 求斐波那契数列的第n个数
	 * @Param: 给定n
	 * @Return: 返回第n个数
	 * @Q:如果想返回前n个数的和怎么写
	 */
	x, y := 0, 1
	for i := 0; i < n; i++ {
		x, y = y, x+y
		//fmt.Printf("%d\t",x)
	}
	return x //y是第n+1个数
}

/*****************Dividing line*****************/

func fibSum(n int) int {
	/**
	 * @Auther:lichang
	 * @Date: 19-12-25 下午4:55
	 * @Description: 返回斐波那契数列的前n个数之和
	 * @Param: 给定数n
	 * @Return: 前n个数之和
	 */
	x, y := 0, 1
	var sum int
	for i := 1; i < n; i++ { //由于第一个数为0，sum可以不加，因此从1开始，略过第一个数的累加
		x, y = y, x+y
		sum += x
	}
	return sum
}

/***********Dividing line***********/
/*
var cwd string

func init(){*/
/**
 * @Author: lichang
 * @date: 19-12-28   20:07
 * @Description: 这段代码有一个隐藏的bug，首先定义了一个包级变量cwd,然后在init函数里又
 * 使用:=重新定义了cwd和err，也就是说，init里的cwd和包级变量cwd其实不是一个变量，全局的cwd
 * 并没有被正确初始化，而且看似正常的日志输出更是会让这个bug更加隐晦
 * @Param:
 * @Return:
 */

/*cwd, err:=os.Getwd()
	if err != nil{
		log.Fatalf("os.Getwd failed : %v\n", err)
	}
	log.Printf("working directory = %s", cwd)
}
*/
