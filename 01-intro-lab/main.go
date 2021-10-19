package main

import (
	"fmt"
	"rsc.io/quote"
	"01-intro-lab/stringutil"
)

func main()	{
	s := "Hello from Golang!"
	for i:=0; i < len(s); i++ {
		fmt.Printf("%#U starts at byte position %d\n", s[i], i)
	}
	fmt.Println(s)
	goquote := quote.Go()
	fmt.Println(goquote)
	fmt.Println(stringutil.Reverse(goquote))
}