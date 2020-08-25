package tempconv_flag

//*celsiusFlag는 flag.Value 인터페이스를 충족한다.

import (
	"flag"
	"fmt"
	"go_practice/ch2/tempconv"
)

type celsiusFlag struct { tempconv.Celsius }

func (f *celsiusFlag) Set(s string) error{
	var unit string
	var value float64

	fmt.Sscanf(s, "%f%s", &value, &unit)

	switch(unit){
	case "C", "℃":
		f.Celsius = tempconv.Celsius(value)
		return nil

	case "F", "℉":
		f.Celsius = tempconv.FToC(tempconv.Fahrenheit(value))
		return nil

	}

	return fmt.Errorf("invalid temperature %q", s)
}

// CelsiusFlag는 CelsiusFlag 를 지정된 시간, 기본값, 사용법으로 정의하고 플래그 변수의 주소를 반환한다.
// 인자 flag 에는 양과 단위가 있어야한다. ex) "100C"
func CelsiusFlag(name string, value tempconv.Celsius, usage string) *tempconv.Celsius{
	f := celsiusFlag{value}
	flag.CommandLine.Var(&f, name, usage)
	return &f.Celsius
}