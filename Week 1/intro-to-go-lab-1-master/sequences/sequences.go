package main

import(
	"fmt"
)

func addOne(a int) int {
	return a + 1
}

func square(a int) int {
	return a * a
}

func double(slice []int) {
	slice = append(slice, slice...)
}

func mapSlice(f func(a int) int, slice []int) {
	for i,v := range slice{
		slice[i] = f(v)
	}
}

func mapArray(f func(a int) int, array [3]int) [3]int {
	for i := 0; i < 3; i++{
		array[i] = f(array[i])
	}
	return array
}

func main() {

	intsSlice := []int{1,2,3,4,5}
	mapSlice(addOne, intsSlice)
	fmt.Println(intsSlice)

	var intsArray [3]int
	intsArray[0] = 1
	intsArray[1] = 2
	intsArray[2] = 3
	intsArray = mapArray(addOne, intsArray)
	fmt.Println(intsArray)

}
