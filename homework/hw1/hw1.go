package hw1

import (
	"container/ring"
)

func FindWinner(n, m int) interface{} {

	peopleCircle := ring.New(n)

	for i := 1; i < peopleCircle.Len() + 1; i++ {
		peopleCircle.Value = i
		peopleCircle = peopleCircle.Next()
	}
	//peopleCircle.Do(func(i interface{}) {
	//	fmt.Print(i.(int)) //  printing on one line
	//})
	//fmt.Println()


	for n != 1{
		peopleCircle = peopleCircle.Move(m-2)
		//peopleCircle.Do(func(i interface{}) {
		//	fmt.Print(i.(int)) //  printing on one line
		//})
		//fmt.Println()

		//peopleCircle = peopleCircle.Prev()
		peopleCircle.Unlink(1)
		peopleCircle = peopleCircle.Move(1)
		//peopleCircle.Do(func(i interface{}) {
		//	fmt.Print(i.(int)) //  printing on one line
		//})
		//fmt.Println()

		n--
	}
	return peopleCircle.Value
}

//func main() {
//
//	return findWinner(8, 3)
//}

// 3, 6, 1, 5, 2, 8, 4

// , , , 4, , , 7,


// 1 ,2 ,3 ,4 , 5, 6, 7, 8