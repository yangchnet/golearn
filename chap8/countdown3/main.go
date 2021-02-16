package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	abort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1)) // read a single byte
		abort <- struct{}{}
	}()

	fmt.Println("commencing countdown. Press return to abort")
	tick := time.Tick(1 * time.Second) // goroutine泄露
	for countdown := 10; countdown > 0; countdown-- {
		fmt.Println(countdown)
		select {
		case <-tick:
			// do nothing
		case <-abort:
			fmt.Println("launch abnroted")
			return
		}
	}
	launch()
}

func launch() {
	fmt.Println("Lift off!")
}
