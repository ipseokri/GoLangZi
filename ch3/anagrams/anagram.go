package anagrams
// anagrams 는 두 문자열이 같은 문자를 반대의 순서로 갖는 아나그램인지 검사하는 함수이다.

import (
	"bytes"
	"fmt"
)

func AnalyzingAnagrams(s1 string, s2 string) bool{
	var buf bytes.Buffer
	for i := len(s1)-1; i >= 0; i --{
		fmt.Fprintf(&buf, "%s", string(s1[i]))
	}

	return buf.String() == s2
}