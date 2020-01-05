package main
/**
 * @Description: 在原有slice内存空间之上返回不包含空字符串的列表
 * @author lichang
 * @date 20-1-4
 * @time 下午8:03
*/

import "fmt"

func nonempty(strings []string) []string{
	i:=0
	for _, s := range strings{
		if s!=""{
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}

func main(){
	data:=[]string{"one", "", "three"}
	fmt.Printf("%q\n", nonempty(data))
	fmt.Printf("%q\n",data)
}
