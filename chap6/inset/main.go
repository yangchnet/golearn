package main

import (
	"bytes"
	"fmt"
)

/**
 * @Description: Bit数组
 * @author lichang
 * @date 20-1-13
 * @time 下午2:41
*/
type InSet struct {
	words []uint64
}

var pc [256]byte = func() (pc [256]byte) {
	/**
	   @Description: 可从pc数组中取出对应的二进制位数
	   @Param:
	        byte数组
	   @Return:
	        对应二进制表示中有几个1
	 */
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
	return
}() //这里的括号表示对这个匿名函数进行调用

func (s *InSet) Len() int  {
	sum := 0
	for _, word := range s.words{
		wordlen := int(pc[byte(word>>(0*8))]+  // 每行取出8bit中的1的个数相加
			pc[byte(word>>(1*8))]+
			pc[byte(word>>(2*8))]+
			pc[byte(word>>(3*8))]+
			pc[byte(word>>(4*8))]+
			pc[byte(word>>(5*8))]+
			pc[byte(word>>(6*8))]+
			pc[byte(word>>(7*8))])
		sum += wordlen
	}
	return sum
}

func (s *InSet) Remove(x int) {
/**
   @Description: 移除某个存在的数字，如果不存在，无操作
 */
	word, bit := x/64, uint(x%64)
	if word > len(s.words){  // 如果不存在，直接返回
		return
	}
	s.words[word] &^= ( 1 << bit)
	//s.words[word] &= ( 0 << bit)
}

func (s *InSet) Copy() *InSet {
	p := InSet{}
	p = *s
	return &p

}

func (s *InSet) Clear()  {
	/**
	   @Description: 清除全部数字
	 */
	for i, _ := range s.words{
		s.words[i] = 0
	}

}

func (s *InSet) Has (x int) bool{
	/**
	   @Description: 判断集合中是否存在某个数
	   @Param:
	        x: 给定数
	   @Return:
	        bool: 是否存在
	 */
	word, bit := x/64, uint(x%64)  // word标记x存储在哪个字里， bit标记x存储在这个字的哪一位
	return word < len(s.words) && s.words[word]&(1<<bit) != 0  //检查字是否存在及字的特定位是否为1
}

func (s *InSet) Add (x int) {
	/**
	   @Description: 向集合中增加一个数
	   @Param:
	        x：给定的数
	   @Return:

	 */
	word, bit := x/64, uint(x%64)
	for word >= len(s.words){   // 如果字的个数不够，则增加一个
		s.words = append(s.words, 0)
	}
	s.words[word] |= (1 << bit) //目标位设为1

}

func (s *InSet) UnionWith(t *InSet){
	/**
	   @Description: 求并集
	   @Param:
			t: 另一个集合
	   @Return:

	 */
	for i, tword := range t.words {
		if i < len (s.words){
			s.words[i] |= tword
		}else {
			s.words = append(s.words, tword)
		}
	}
}

func (s *InSet) String() string{
	/**
	   @Description: 将集合按字符串格式输出
	   @Param:

	   @Return:
	        字符串
	 */
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words{
		if word == 0 {
			continue
		}
		for j := 0 ;j < 64;j++ {
			if word&(1<<uint(j)) != 0{
				if buf.Len() > len("{"){
					buf.WriteByte(',')
				}
				_, _ = fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

func main() {
	var x ,y InSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	x.Add(8)
	fmt.Println(x.Len())
	fmt.Println(x.String())
	x.Remove(1)
	//s := x.Copy()
	fmt.Println(x.Len())
	fmt.Println(x.String())
	//s.Add(99)
	//fmt.Println(s.Len())
	//fmt.Println(s.String())

	y.Add(9)
	//y.Add(42)
	//fmt.Println(y.Len())
	//fmt.Println(y.String())

	//
	//x.UnionWith(&y)
	//fmt.Println(x.String())
	//
	//fmt.Println(x.Has(9), x.Has(123))
}

