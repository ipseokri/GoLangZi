package main

import "fmt"

func main(){
	p, q := 0, 1

	var n int = 6
	var i int = 0

	for i = 0; i < n-1; i++{
		p, q = q, p+q
	}

	fmt.Println(n, "번째 피보나치 수 : ", p)



}

