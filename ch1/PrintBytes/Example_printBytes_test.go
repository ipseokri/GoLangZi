package PrintBytes

import "fmt"

func Example_printBytes(){
	s := "가나다"
	for i := 0; i < len(s); i ++{
		fmt.Printf("%x:", s[i])
	}
	fmt.Printf("%x\n",s)
	fmt.Printf("% x\n",s)
	fmt.Println()

	b := []byte("가나다")
	b[2]++
	fmt.Println(string(b))
	// Output:
	// ea:b0:80:eb:82:98:eb:8b:a4:eab080eb8298eb8ba4
	// ea b0 80 eb 82 98 eb 8b a4
	//
	// 각나다
}

