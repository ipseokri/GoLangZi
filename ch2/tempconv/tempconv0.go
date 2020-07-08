// tempconv package is calculate Celsius to Fahrenheit or calculate Celsius to Kelvin
package tempconv

import "fmt"

type Celsius float64
type Fahrenheit float64
type Kelvin float64

const(
	absoluteZeroC Celsius = -273.15
	FreezingC Celsius = 0
	BoilingC Celsius = 100

	ZeroKelvin Celsius = -273.15

)

func (c Celsius) String() string {return fmt.Sprintf("%g℃",c)}
func (f Fahrenheit) String() string{return fmt.Sprintf("%g℉",f)}
func (k Kelvin) String() string{return fmt.Sprintf("%gK",k)}


