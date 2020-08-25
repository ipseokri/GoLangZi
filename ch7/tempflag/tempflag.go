package main

import (
	"flag"
	"fmt"
	"go_practice/ch7/tempconv"
)

var temp = tempconv_flag.CelsiusFlag("temp", 20.0, "the temperature")

func main(){
	flag.Parse()
	fmt.Println(*temp)
}
