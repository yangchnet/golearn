package main

import (
	"crypto/sha256"
	"fmt"
)

/**
 * @Description: 计算两个SHA256哈希码中不同的bit数目
 * @author lichang
 * @date 20-1-3 
 * @time 下午3:17
*/

var pc [256]byte

func main(){
	x1 := sha256.Sum256([]byte("x"))
	x2 := sha256.Sum256([]byte("X"))
	fmt.Printf("%x\n", x1)
	fmt.Printf("%x\n", x2)
	//fmt.Printf("%T\t%T", x1, x2)
	difference := ShaCompare(x1, x2)
	fmt.Printf("Two sha256 value between 'x' and 'X' have %d different bit\n", difference)

}

func init(){
	for i := range pc{
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func ShaCompare(x1, x2 [32]uint8) int {
	/**
	 * @Auther:lichang
	 * @Date: 20-1-3 下午3:21
	 * @Description: 统计两个哈希值中不同bit值的数目
	 * @Param: 两个uint[32]数组
	 * @Return:
	 */
	sum := 0
	for i:=0;i<32;i++{
		//y1 = fmt.Sprintf("")
		sum += int(pc[x1[i]^x2[i]])
	}
	return sum
}




