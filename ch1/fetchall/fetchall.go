// fetchall은 URL을 병렬로 반입하고 시간과 크기를 보고한다.
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main(){
	start:= time.Now()
	ch := make(chan string)
	for _,url := range os.Args[1:]{
		go fetch(url, ch) // 고루틴 시작
	}
	for range os.Args[1:]{
		fmt.Println(<-ch) // ch 채널에서 수신
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())

}

func fetch(url string, ch chan <- string){
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil{
		ch <- fmt.Sprint(err)
		return
	}

	nbytes,err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
	if err!= nil{
		ch <- fmt.Sprintf("while reading %s : %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%2.fs %7d %s", secs, nbytes, url)

}
