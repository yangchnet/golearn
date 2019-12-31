package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	input := bufio.NewScanner(os.Stdin)
	for input.Scan(){
		fmt.Println(comma(input.Text()))
	}
}

func comma(s string) string{
	n := len(s)
	if n<=3{
		return s
	}
	return s[:3] + "," + comma(s[3:])
}