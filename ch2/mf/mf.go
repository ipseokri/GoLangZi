// mf is change the factor to meter or feet
package main

import (
	"awesomeProject1/ch2/heigthconv"
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

		f := heigthconv.Feet(t)
		m := heigthconv.Meter(t)

		fmt.Printf("%s = %s, %s = %s \n", f, heigthconv.FToM(f), m, heigthconv.MToF(m))
	}
}