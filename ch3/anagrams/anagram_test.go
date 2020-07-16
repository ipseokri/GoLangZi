package anagrams

import "fmt"

func ExampleAnalyzingAnagrams() {
	fmt.Println(AnalyzingAnagrams("abc!efg", "gfe!cba"))
	// OUTPUT:
	// true
}
