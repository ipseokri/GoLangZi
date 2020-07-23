//Findlinks1 은 표준 입력에서 읽어드린 HTML 문서 안의 링크를 출력한다.
package main

import(
	"fmt"
	"os"
	"github.com/mjibson/goread/_third_party/golang.org/x/net/html"

)

func main(){
	doc, err := html.Parse(os.Stdin)
	if err != nil{
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	for _, link := range visit(nil, doc){
		fmt.Println(link)
	}
}

// visit는 n 에서 찾은 각 링크를 links에 추가하고 결과를 반환한다.
func visit(links []string, n *html.Node) []string{
	if n.Type == html.ElementNode && n.Data == "a"{
		for _,a := range n.Attr{
			if a.Key == "href"{
				links = append(links, a.Val)
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}

	return links
}