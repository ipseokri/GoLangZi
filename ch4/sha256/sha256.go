package main

import (
	"crypto/sha256"
	"fmt"
)

func main(){
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))

	fmt.Printf("% x\n% x\n%t\n%T\n", c1,c2,c1==c2, c2)

	for i := 0;  i < 8; i ++{
		fmt.Printf("i : %x \n", c1[i])
		fmt.Printf("i2 : %x \n", c2[i])

	}
}
