package main

import "fmt"

func search(data []int, value int) bool {

	size := len(data)

	var mid int

	low := 0
	high := size - 1

	for low <= high {
	
		mid = low + (high-low)/2
	      if data[mid] == value {
	        return true
	       } else {
	      if data[mid] < value {
	         low = mid + 1
	         } else {
	         high = mid - 1
	       }
	    }
	}
	return false
	}

func main(){
    items := []int{1,2, 9, 20, 31, 45, 63, 70, 100}
    fmt.Println(search(items, 60))
}