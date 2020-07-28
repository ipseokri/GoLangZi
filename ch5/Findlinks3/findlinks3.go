package main

import (
	"fmt"
	"go_practice/ch5/links"
	"log"
	"os"
)

// breadthFirst는 작업 목록의 각 항목으로 f를 호출한다.
// f 에서 반환된 항목은 worklist에 추가된다.
// f 는 한 항목에 최소한 한번 실행된다.
func breadthFirst(f func(item string) []string, worklist []string){
	seen := make(map[string]bool)
	for len(worklist) > 0{
		items := worklist
		worklist = nil
		for _, item := range items{
			if !seen[item]{
				seen[item] = true
				worklist = append(worklist, f(item)...)
			}
		}
	}
}

func crawl(url string) []string{
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil{
		log.Print(err)
	}
	return list
}

func main(){
	// 웹을 깊이 우선으로 탐색하며, 커맨드라인 인수로 시작한다.
	breadthFirst(crawl, os.Args[1:])
}