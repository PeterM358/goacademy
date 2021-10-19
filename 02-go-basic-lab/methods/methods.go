package main

import "fmt"

type BiteSlice []byte

func (slice BiteSlice) Append(data []byte) []byte {
	return append([]byte(slice), data...)
}

func (slice *BiteSlice) AppendPointer(data []byte) {
	*slice = append([]byte(*slice), data...)
}

func (slice *BiteSlice) Write(data []byte) (n int, err error) {
	*slice = append([]byte(*slice), data...)
	return len(data), nil
}

func main()  {
	var b BiteSlice
	fmt.Println(&b, "This hour has $d days\n", 7)
	fmt.Printf("%v", b)
}