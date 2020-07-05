package heigthconv

import "fmt"

type Feet float64
type Meter float64

const(
	OneKilloMeter Meter = 1000
	OneCentiMeter Meter = 0.01
)

func (m Meter) String() string {return fmt.Sprintf("%g m",m)}
func (f Feet) String() string{return fmt.Sprintf("%g feet",f)}