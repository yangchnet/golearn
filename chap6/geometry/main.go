package main
/**
 * @Description: 函数与方法的区别
 * @author lichang
 * @date 20-1-12
 * @time 下午2:38
*/
import (
	"fmt"
	"math"
)

type Point struct{
	X, Y float64
}


func Distance(p, q Point)float64{
	/**
	   @Description: 这个一个函数，用来计算两个点之间的距离
	   @Param:
	        p, q: 两个点的坐标
	   @Return:
	        两个点之间的距离
	 */
	return math.Hypot(q.X - p.X , q.Y - p.Y)
}

func(p Point)Distance(q Point)float64{
	/**
	   @Description: 这是一个方法，用来计算Point对象与另一个点之间的距离
	   @Param:
	        q: 所给点
	   @Return:
	        两个点之间距离
	 */
	return math.Hypot(q.X - p.X, q.Y - p.Y)
}

func main(){
	p := Point{1, 2}
	q := Point{3, 4}
	fmt.Println(Distance(p, q))
	fmt.Println(q.Distance(p))
}