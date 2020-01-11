// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 146.

// The trace program uses defer to add entry/exit diagnostics to a function.
package main

import (
	"fmt"
	"log"
	"time"
)

//!+main
func bigSlowOperation() {
	/**
	   @Description: trace 返回一个函数值，这个函数值只会在bigSlowOperation函数退出时被调用，
	而trace函数本身，在defer函数被定义时就已经被调用
	   @Param:

	   @Return:

	 */
	defer trace("bigSlowOperation")() // don't forget the extra parentheses
	// ...lots of work...
	fmt.Println("hello")
	time.Sleep(10 * time.Second) // simulate slow operation by sleeping
}

func trace(msg string) func() {
	start := time.Now()
	log.Printf("enter %s", msg)
	return func() { log.Printf("exit %s (%s)", msg, time.Since(start)) }
}

//!-main

func main() {
	bigSlowOperation()
}

/*
!+output
$ go build gopl.io/ch5/trace
$ ./trace
2015/11/18 09:53:26 enter bigSlowOperation
2015/11/18 09:53:36 exit bigSlowOperation (10.000589217s)
!-output
*/
