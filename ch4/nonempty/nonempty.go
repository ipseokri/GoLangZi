package main

import "fmt"

// Nonempty는 슬라이스 직접 변경 알고리즘의 예이다

// nonempty는 비어있지 않은 문자열을 갖는 슬라이스만 반환한다.
// 내부 배열은 호출에서 변경된다.
func nonempty(strings []string) []string{
	i := 0
	for _,s := range strings{
		if s != ""{
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}

func main(){
	data := []string{"one", "", "three"}
	fmt.Printf("%q\n", nonempty(data))//'["one", "three"]'
	fmt.Printf("%q\n", data) // '["one", "three", "three"
	// ]'
}