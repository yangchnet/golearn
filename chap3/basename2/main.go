package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main(){
	input := bufio.NewScanner(os.Stdin)
	a := input.Scan()
	b := 1
	b ++
	fmt.Println(a)
	for input.Scan(){
		fmt.Println(basename(input.Text()))
	}
}

func basename(s string)string {
	slash := strings.LastIndex(s, "/")
	s = s[slash+1:]
	if dot := strings.LastIndex(s, ".") ;dot >= 0{
		s = s[:dot]
	}
	return s

}
