package tempconv

//CToF is translate Celsius degree to Fahrenheit degree
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5+32)}

//FToC is translate Fahrenheit degree to Celsius degree
func FToC(f Fahrenheit) Celsius { return Celsius((f-32)*5/9)}