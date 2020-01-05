package main
/**
 * @Description: 移除某个位置上的元素
 * @author lichang
 * @date 20-1-4 
 * @time 下午8:07
*/
import "fmt"

func remove(slice []int, i int) []int {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}

func main(){
	s := []int{5,6,7,8,9}
	fmt.Println(remove(s, 2))
}