package basename1

import "strings"

//moresimple version against basename 1

func basename2 (s string) string{
	slash := strings.LastIndex(s, "/") // "/"가 없으면 -1
	s = s[slash+1:]
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}

	return s
}