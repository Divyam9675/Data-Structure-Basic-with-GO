package main

import "fmt"

func RotateArray(data []int, k int) {
	
	n := len(data)
	ReverseArray(data, 0, k-1)
	ReverseArray(data, k, n-1)
	ReverseArray(data, 0, n-1)
	
}
	
	func ReverseArray(data []int, start int, end int) {
	
		i := start
		j := end
	     for i < j {

		 data[i], data[j] = data[j], data[i]

	        i++
			j--
		}
		printArray(data, end)
	}

	
	func printArray(data []int, n int){ 
		
			fmt.Println(data)
		
			
	} 
		 

	func main(){
		items := []int{1,2, 9, 20, 31, 45, 63, 70, 100}

		 RotateArray(items, 5)

		 


	//	fmt.Println(items[])
	}	


