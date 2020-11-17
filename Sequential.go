package main 

import "fmt"

func SequentialSearch(data []int, value int) bool {

	size := len(data)

	for i := 0; i < size; i++ {

		if value == data[i] {
	         return true

			}
	}
	return false
}
	

func main(){

	var x = []int{10,20,30}

//	for y:=0; y<15; y++{

//	z :=	append(x, y)

	fmt.Println(x)

	fmt.Println(SequentialSearch(x, 20 ))
		
}