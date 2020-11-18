package main

import "fmt"

func towerOfHanoi(n int, from_rod string, to_rod string, aux_rod string){  

	if n == 1 {  
          fmt.Println(" Move disk 1 from rod ", from_rod , " to rod " , to_rod)  
         return;  
     }  

	 towerOfHanoi(n - 1, from_rod, aux_rod, to_rod);  

	 fmt.Println( " Move disk ", n ," from rod ", from_rod, " to rod ", to_rod)  

	 towerOfHanoi(n - 1, aux_rod, to_rod, from_rod);  

	 }  

// Driver code 
func main(){  

	n := 4;   

	 towerOfHanoi(n, "A", "C", "B") // A, B and C are names of rods  
	 
}  