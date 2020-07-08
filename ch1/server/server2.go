//server2는 최소한의 "echo" 및 카운터 서버다.
package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var mu sync.Mutex
var count int

func main(){
	http.HandleFunc("/", Handler)
	http.HandleFunc("/count", Counter)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
//Handler 는 요청된 URL의 Path 구성 요소를 반환한다.
func Handler(w http.ResponseWriter, r * http.Request){
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}
//Counter는 지금까지 요청된 수를 반환한다.
func Counter(w http.ResponseWriter, r * http.Request){
	mu.Lock()
	fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
}
