package main


import (
	"fmt"
	"os"
)

func main(){
	var s, sep string
	for i := 0; i < len(os.Args); i++{
		s+= sep+os.Args[i]
		sep = " "
	}
	fmt.Println("output: ", s)

	for i, j := range os.Args[0: ]{
		fmt.Println("Output : ",i, ",", j)
	}

}
