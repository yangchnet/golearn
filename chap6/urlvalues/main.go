package main

import "fmt"

//定义了一个map类型，其key是string类型，value是string数组
type Values map[string][]string

func (v Values)Get(key string)string{
	/**
	   @Description: 返回v的key对应的value的第一个值
	   @Param:
	        key： map中的value值
	   @Return:
	        对应key的value中第一个值
	 */
	if vs := v[key]; len(vs) > 0 {
		return vs[0]
	}
	return ""
}

func (v Values) Add(key, value string){
	/**
	   @Description: Add values to key.It append to any existing values associated with key.
	   @Param:
	        key:
			value:
	   @Return:

	 */
	v[key] = append(v[key], value)
}

func main(){
	m := Values{"lang":{"en"}}
	m.Add("item", "1")
	m.Add("item", "2")

	fmt.Println(m.Get("lang"))
	fmt.Println(m.Get("q"))
	fmt.Println(m.Get("item"))
	fmt.Println(m["item"])

	m = nil
	fmt.Println(m.Get("item"))
	//Values(nil).Add("item", "3")
}
