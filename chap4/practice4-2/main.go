package main
/**
 * @Description: 程序默认打印标准输入的以sha256哈希码，也可以通过命令行标准参数选择sha384或sha512哈希算法
 * @author lichang
 * @date 20-1-3
 * @time 下午4:42
*/
import (
	"bufio"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"os"
)

func main(){
	input := bufio.NewScanner(os.Stdin)
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "sha384":
			{
				for input.Scan() {
					fmt.Printf("%x\n", sha512.Sum384([]byte(input.Text())))
				}
			}
		case "sha512":
			{
				for input.Scan() {
					fmt.Printf("%x\n", sha512.Sum512([]byte(input.Text())))
				}
			}
		}
	}else{
		for input.Scan(){
			fmt.Printf("%x\n", sha256.Sum256([]byte(input.Text())))
		}
	}
}