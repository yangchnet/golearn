package main

/**
 * @Description: 非递归版本的comma
 * @author lichang
 * @date 19-12-31
 * @time 下午10:10
 */
import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		fmt.Println(commaNoRecursive(input.Text()))
	}
}
func commaNoRecursive(s string) string {
	var buf bytes.Buffer
	for index, value := range (s) {
		if (index+1)%3 == 0 && index != 0 {
			_, _ = fmt.Fprintf(&buf, "%c", value)
			_ = buf.WriteByte(',')
		} else {
			_, _ = fmt.Fprintf(&buf, "%c", value)

		}
	}
	return buf.String()
}
