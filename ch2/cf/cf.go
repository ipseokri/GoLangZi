// cf is change the factor to Celsius or Fahrenheit
package main

import (
	"go_practice/ch2/tempconv"
	"fmt"
	"os"
	"strconv"
)

func main(){
	for _,arg := range os.Args[1:]{
		t, err := strconv.ParseFloat(arg,64)
		if err != nil{
			fmt.Fprint(os.Stderr, "cf : %v\n", err)
			os.Exit(1)
		}

		f := tempconv.Fahrenheit(t)
		c := tempconv.Celsius(t)

		fmt.Printf("%s = %s, %s = %s \n", f, tempconv.FToC(f), c, tempconv.CToF(c))
	}
}