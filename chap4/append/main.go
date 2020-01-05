package main

import "fmt"

/**
 * @Description: 了解append函数的内部作用原理
 * @author lichang
 * @date 20-1-4
 * @time 下午7:19
*/

func appendInt(x []int, y int) []int{  
	/**
	   @Description: 这里[]int是一个类型，如果传入[4]int则会产生错误
	   @Param：
			x: 一个[]int类型数据，可以传入切片
			y: int整数
	   @Return:
	 */
	var z []int
	zlen:=len(x) + 1
	if zlen <= cap(x){
		z = x[:zlen]
	}else{					//未来增加后的长度>当前容量
		zcap:=zlen
		if zcap<2*len(x){	//增加一倍
			zcap=2*len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x)
	}
	z[len(x)] = y
	return z
}

func main(){
	var x, y []int
	for i:=0;i<10;i++{
		y = appendInt(x, i)
		fmt.Printf("%d cap = %d \t%v\n", i, cap(y), y)
		x = y
	}

	/*****************Dividing line*****************/
	var x1 = [4]int{1,2,3,4}
	var y1 = 5
	c := appendInt(x1[:], y1)
	fmt.Printf("%v", c)
}
