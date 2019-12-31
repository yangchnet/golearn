package main
/**
 * @Description: 从一个文件链接中提取出其基本文件名（不包含前缀和后缀）
 * @author lichang
 * @date 19-12-31
 * @time 下午4:55
*/

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		if input.Text() == "-1"{
			break
		}
		fmt.Println(basename(input.Text()))
	}
}

func basename(s string) string {
	/**
	 * @Auther:lichang
	 * @Date: 19-12-31 下午4:57
	 * @Description:
	 * @Param: 文件链接
	 * @Return: 文件名
	 */
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
		}
	}
	return s
}
