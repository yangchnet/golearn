package popcount

/**
 * @Description:
 * @author:lichang
 * @date: 19-12-28
 * @time: 19:20
 */

import "fmt"

// var pc [256]byte

// func init(){
// 	/**
// 	 * @Author: lichang
// 	 * @date: 19-12-28   19:37
// 	 * @Description: 构造出辅助表格pc,pc的每一项对应的值为其下标值对应二进制中1的个数
// 	 * @Param: 无
// 	 * @Return: pc[]
// 	 */

// 	for i:=range pc{
// 		pc[i] = pc[i/2] + byte(i&1)	//byte是uint8的别名
// 	}
// }

/***********Dividing line***********/

var pc [256]byte = func() (pc [256]byte) {
	/**
	 * @Author: lichang
	 * @date: 19-12-28   19:48
	 * @Description:  init函数的另一种写法，匿名函数
	 * @Param:
	 * @Return:
	 */

	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
	return
}() //这里的括号表示对这个匿名函数进行调用

/***********Dividing line***********/

func PopCount(x uint64) int {
	/**
	 * @Author: lichang
	 * @date: 19-12-28   19:19
	 * @Description: 统计一个数字的二进制表示有几位是1
	 * @Param: 整数
	 * @Return: 整数的二进制表示有几个二进制1
	 */

	fmt.Printf("hello")
	for i := range pc {
		fmt.Printf("%v", pc[i])
	}
	fmt.Printf("\n")
	// 每次取x二进制表示的8位，可从pc中直接取出对应的二进制1个数
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}
