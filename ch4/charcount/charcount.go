package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

// charcount는 유니코드 문자의 카운트를 계산한다.
// 아직 실행 못시킴, 실행 방법을 강구할 것


func main(){
	counts := make(map[rune]int) // 유니코드 문자 카운트
	var utflen [utf8.UTFMax+1]int // UTF-8 인코딩 길이 카운트
	invalid := 0 // 잘못된 UTF - 8 문자 카운트

	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune() // rune, nbytes, error 반환
		if err == io.EOF {
			break
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid ++
			continue
		} else{
			break
		}
		counts[r] ++
		utflen[n] ++
 	}

 	fmt.Printf("rune\tcount\n")
	for c, n := range counts{
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Print("\nlen\tcount\n")
	for i, n := range utflen{
		if i > 0 {
			fmt.Printf("%d\t%d\n", i , n)
		}
	}
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}

}