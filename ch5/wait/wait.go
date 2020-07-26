package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

//WaitForServer는 URL 로 지정된 서버로 접속을 시도한다.
// 1분간 지수 단위로 백호프(exponential back-off)를 수행한다.
// 모든 시도가 실패하면 오류를 보고한다.

func WaitForServer(url string) error{
	const timeout = 1 * time.Minute
	deadline := time.Now().Add(timeout)
	for tries := 0; time.Now().Before(deadline); tries ++{
		_, err := http.Head(url)
		if err == nil{
			return nil // 성공
		}
		log.Printf("server not responding (%s); retrying...", err)
		time.Sleep(time.Second << uint(tries)) // 지수 단위 백오프
	}
	return fmt.Errorf("server %s failed to respond after %s", url, timeout)
}
