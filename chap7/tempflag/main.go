package main

import (
	"flag"
	"fmt"
	"golearn/chap7/tempconv"
)

var temp = tempconv.CelsiusFlag("temp", 20.0, "the temperature")
//var ftemp = tempconv.FahrenheitFlag("temp", 100.0, "the temperature")

func main(){
	flag.Parse()
	fmt.Println(*temp)
}
