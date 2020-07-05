package main

import "fmt"


func main(){

	var n = 5
	var a = 1

	for n > 0 {
		fmt.Println("타잔이", 10*a, "원짜리 팬티를 입고, ", 10*(a+1), "원짜리 칼을 차고 노래를 한다.")
		n--
		a++
	}
}
