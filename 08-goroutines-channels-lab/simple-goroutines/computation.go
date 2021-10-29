package main

import (
	"fmt"
	"time"
)

func  main()  {
	compute("Long computation")
}

func compute(msg string)  {
	for i := 0; ; i++ {
		fmt.Println(msg, i)
		time.Sleep(1000)
	}
}