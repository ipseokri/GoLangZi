package main

import "fmt"

func appendInt(x []int, y int) [] int{
	var z []int
	zlen := len(x)+1
	if zlen <= cap(x){
		// 확장할 공간이 있다. 슬라이스를 확장하라.
		z = x[:zlen]
	}else{
		// 공간이 부족하다. 새 배열을 할당하라
		// 복잡도의 선형 증가를 막기 위해 도배로 증가시킨다.
		zcap := zlen
		if zcap < 2*len(x){
			zcap = 2 * len(x)
		}


		z = make([]int, zlen, zcap)
		copy(z,x)
	}
	z[len(x)] = y
	return z
}

func main(){
	var x, y []int
	for i := 0; i < 10; i ++{
		y = appendInt(x, i)
		fmt.Printf("%d cap=%d\t%v\n", i, cap(y), y)
		x = y
	}
}
