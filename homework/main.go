package main

import (
	"fmt"
	"homework/hw1"
)


func main() {
 	var n, m int
	 fmt.Scanln(&n)
	 fmt.Scanln(&m)
	fmt.Println(hw1.FindWinner(n, m))
}
