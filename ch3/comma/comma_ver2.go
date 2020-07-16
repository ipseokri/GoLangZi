package comma

import (
	"bytes"
	"fmt"
)

// comma 는 양의 10진 정수 문자열에 콤마를 삽입한다.

func CommaVer2(s string) string{
	var buf bytes.Buffer
	for i, v  := range s {
		if i % 3 == 0 {
			if i != 0 {
				buf.WriteString(",")
			}
		}
		fmt.Fprintf(&buf , "%s" , string(v))
	}

	fmt.Println(buf.String())

	return buf.String()
}