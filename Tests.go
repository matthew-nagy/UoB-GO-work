package main

import(
	"fmt"
)

func gitFunc() func()int{
	num := 0
	return func()int{
		num = num + 1
		return num
	}
}

func main(){

	inc := gitFunc()
	inc()
	inc()
	fmt.Println(inc())

	inc2 := gitFunc()
	fmt.Println(inc2())
	fmt.Println(inc())

}