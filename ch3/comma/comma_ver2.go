package comma

import (
	"bytes"
	"fmt"
)

// comma 는 양의 10진 정수 문자열에 콤마를 삽입한다.

func comma_ver2(s string) string{
	s = "12345"
	var buf bytes.Buffer
	for i, v  := range s {
		if i % 3 == 0 {
			buf.WriteString(", ")
		}
		fmt.Fprintf(&buf , "%s" , string(v))
	}

	fmt.Println(buf.String())

	// Output:
	// 12,345
	return buf.String()
}