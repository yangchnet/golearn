package main
/**
 * @Description: 此代码还存在错误
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

func main(){
	input := bufio.NewScanner(os.Stdin)
	for input.Scan(){
		fmt.Println(commaNoRecursive(input.Text()))
	}
}
func commaNoRecursive(s string) string{
	var buf  bytes.Buffer
	for index, value := range(s){
		if (index+1) % 3 ==0{
			_ = buf.WriteByte(',')
		}else{
			_, _ = fmt.Fprintf(&buf, "%v", value)

		}
	}
	return buf.String()
}
