package main

import "fmt"

func equal(x, y map[string]int)bool{
	/**
	   @Description: 判断两个字典是否相同
	   @Param:
	        两个字典
	   @Return:
	   		布尔型
	 */
	if len(x) != len(y){
		return false
	}
	for k, xv := range x{
		if yv, ok := y[k]; !ok || yv != xv{
			return false
		}
	}
	return true
}

func main(){
	d1 := make(map[string]int)
	d2 := make(map[string]int)
	d1["one"] = 1
	d2["one"] = 1
	e := equal(d1, d2)
	fmt.Printf("%v", e)
}