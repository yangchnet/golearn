package main

import "fmt"

func counter(out chan<- int){
	/**
	    @Description:chan<-是一个只用来接收的channel
	 */
	for x := 0; x < 100; x++{
		out <- x
	}
	close(out)
}

func squarer(out chan<- int, in <-chan int){
	/**
	    @Des: <-chan是一个只用来发送的channnel
	*/
	for v:= range in {
		out <- v * v
	}
}

func printer(in <-chan int){
	for v:= range in {
		fmt.Println(v)
	}
}

func main(){
	naturals := make(chan int)
	squares := make(chan int)
	go counter(naturals)
	go squarer(squares, naturals)
	printer(squares)
}