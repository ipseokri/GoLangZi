package main

import (
	"fmt"
	"github.com/mjibson/goread/_third_party/golang.org/x/net/html"
	"os"
)

func main(){
	doc, err := html.Parse(os.Stdin)
	if err != nil{
		fmt.Fprintf(os.Stderr, "outline : %v\n", err)
		os.Exit(1)
	}
	outline(nil, doc)
}

func outline(stack []string, n *html.Node){
	if n.Type == html.ElementNode{
		stack = append(stack, n.Data)// 태그 push
		fmt.Println(stack)
	}

	for c:= n.FirstChild; c!= nil; c = c.NextSibling{
		outline(stack, c)
	}
}
