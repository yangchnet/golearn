package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	/**
	   @Description: 忽略map的value而将其当做一个集合，只包含相异的字符串
	   @Param:

	   @Return:

	 */
	seen := make(map[string]bool)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan(){
		line := input.Text()
		if !seen[line]{
			seen[line] = true
			fmt.Println(line)
		}
	}
	if err:=input.Err();err != nil{
		fmt.Fprintf(os.Stderr, "dedup: %v\n", err)
		os.Exit(1)
	}
}
