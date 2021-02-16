package main

import (
	"fmt"
	"time"
)

func main() {
	/**
	  @Destription:利用两个channel连接三个goroutine
	  @Author:lichang
	  @Date:2020/1/19
	  @Time:下午7:33
	*/
	naturals := make(chan int)
	squares := make(chan int)

	//Counter
	go func() {
		for x := 0; ; x++ {
			naturals <- x
			time.Sleep(time.Second)
		}
	}()

	// Squarer
	go func() {
		for {
			x, ok := <-naturals
			if !ok {
				break
			}
			squares <- x * x
		}
	}()

	//Printer
	for {
		fmt.Println(<-squares)
	}
}
