// echo4 is print out commandline variables
package main

import(
	"flag"
	"fmt"
	"strings"
)

var n = flag.Bool("n", false, "omit trailling newline")
var sep = flag.String("s", " ", "separator")

func main(){
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))

	if !*n{
		fmt.Println()
	}
}
