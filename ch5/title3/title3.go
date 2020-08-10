package main

import (
	"fmt"
	"golang.org/x/net/html"
)

//SoleTitle은 비어있지않은 첫번째 title 원소의 텍스트를 반환하며,
// 한개가 아니라면 error를 반환한다.

func soleTitle(doc *html.Node) (title string, err error){
	type bailout struct {}

	defer func() {
		switch p := recover(); p {
		case nil:
			//패닉이 아님

		case bailout{}:
			//"예상된 패닉"
			err = fmt.Errorf("mutiple title elements")

		default:
			panic(p)
		}
	}()

	// 비어있지 않은 title 이 두개 이상이면 재귀에서 탈출한다.
	forEachNode(doc, func(n * html.Node) {
		if n.Type == html.ElementNode && n.Data == "title" && n.FirstChild != nil {
			if title != "" {
				panic(bailout{}) // 여러 title 원소
			}
			title = n.FirstChild.Data
		}
	}, nil)
	if title == ""{
		return "", fmt.Errorf("no title element")
	}
	return title, nil
}


func forEachNode(n *html.Node, pre, post func(n *html.Node)){
	if pre != nil{
		pre(n)
	}
	for c:= n.FirstChild ; c != nil; c = c.NextSibling{
		forEachNode(c, pre, post)
	}
	if post != nil{
		post(n)
	}
}