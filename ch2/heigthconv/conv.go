package heigthconv

// MToF is translate Meter to Feet
func MToF(M Meter) Feet { return Feet(M * Meter(3.28)) }

// FToM is translate Feet to Meter
func FToM(F Feet) Meter { return Meter(F / Feet(3.28)) }