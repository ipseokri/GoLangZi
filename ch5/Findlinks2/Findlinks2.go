package main

import (
	"fmt"
	"github.com/mjibson/goread/_third_party/golang.org/x/net/html"
	"net/http"
	"os"
)

func main(){
	for _, url := range os.Args[1:]{
		links, err := findLinks(url)
		if err != nil{
			fmt.Fprintf(os.Stderr, "findlinks2 : %v\n", err)
			continue
		}
		for _, link := range links{
			fmt.Println(link)
		}
	}
}

// findLinks는 url 에 HTTP GET 요청을 수행하고 결과를 HTML로 파싱한 후 링크를 추출하고 반환한다.
func findLinks(url string) ([]string, error){
	resp, err := http.Get(url)
	if err != nil{
		return nil, err
	}
	if resp.StatusCode != http.StatusOK{
		resp.Body.Close()
		return nil, fmt.Errorf("getting %s : %s", url, resp.Status)
	}

	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil{
		return nil, fmt.Errorf("parsing %s as HTML : %v", url, err)
	}

	return visit(nil, doc), nil

}

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
