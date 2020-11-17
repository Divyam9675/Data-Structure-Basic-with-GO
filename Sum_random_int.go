
// Write a method that will return the sum of the elements of the integer list, given list as an  input argument

package main

import "fmt"

 func create(data []int) int {

	size := len(data)
	total := 0
	
	for index := 0; index < size; index++ {
	
		total = total + data[index]
	
    	}
	
		return total

	}

	func main() {
		var s []int
		for i := 1; i <= 17; i++ {
		s = append(s, i)
		create(s)
		fmt.Printf("No is %d \n", s)
		}
		}
	

