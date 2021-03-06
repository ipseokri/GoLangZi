package main

import (
	"fmt"
	"go_practice/ch6/geometry"
)

func main(){
	perim := geometry.Path{
		{1,1},
		{5,1},
		{5,4},
		{1,1},
	}

	fmt.Println(geometry.Path.Distance(perim))
	fmt.Println(perim.Distance())
}