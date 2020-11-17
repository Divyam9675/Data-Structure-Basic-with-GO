package main

import "fmt"

func search(rect []int , value int)bool{

	

	for _ , item := range rect{

		if  item == value {

			return true

		}

		
	}	

	return false
		
	
}

func main(){

	x := []int{10,20,30,40,50}

	fmt.Println(search(x, 40))

}