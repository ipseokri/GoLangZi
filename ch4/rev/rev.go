package rev
// reverse는 int 슬라이스를 직접 반전시킨다

func reverse(s []int){
	for i, j := 0, len(s)-1; i < j; i , j = i+1, j-1{
		s[i], s[j] = s[j], s[i]
	}
}
