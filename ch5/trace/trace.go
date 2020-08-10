package main

import (
	"log"
	"time"
)

func bigSlowOperation(){
	defer trace("bigSlowOperation")() // 추가 괄호를 잊지 말것
	// ... 다수의 작업 ...
	time.Sleep(10 * time.Second)
}

func trace(msg string) func(){
	strat := time.Now()
	log.Printf("enter %s", msg)
	return func() { log.Printf("exit %s (%s)", msg, time.Since(strat))}
}

func main(){
	bigSlowOperation()
}