package popcount

// pc[i]'s i is population
var pc[256] byte

func init(){
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCount(x uint64) int{
	var total byte = 0

	for i := 0;  i < 8; i ++{
		total += pc[byte(x>>(i*8))]
	}

	return int(total)
}

