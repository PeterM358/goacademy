package main

import (
	"04-struct-methods-lab/mapcopy"
	"fmt"
)

type  Vertex struct {
	Lat, Long float64
}

func main() {
	//m := make(map[string]Vertex)
	//m["Bell Labs"] = Vertex{40.68433, -74.39967}
	//for k, v := range m	{
	//	fmt.Printf("%s -> %v\n", k, v)
	//}
	//
	//m2 := make(map[string]Vertex)
	//m2["Google"] = Vertex{40.68433, -74.39967}
	//for k, v := range m2	{
	//	fmt.Printf("%s -> %v\n", k, v)
	//}
	//fmt.Printf("%s -> %v\n", "Bell Labs", m["Bell Labs"])


	var m mapcopy.GenericMap = make(mapcopy.GenericMap)
	m["Bell Labs"] = Vertex{40.68433, -74.39967}
	//for k, v := range m	{
	//	fmt.Printf("%s -> %v\n", k, v)
	//}

	m2 := mapcopy.GenericMap{}
	m2["Google"] = Vertex{40.68433, -74.39967}
	for k, v := range m2	{
		fmt.Printf("%s -> %v\n", k, v)
	}

	mapcopy.Copymap(m, m2)
}
