package main

import "fmt"


type queue struct{

	item[] int

}



func (l *queue) enque(i int){

	l.item = append(l.item, i)

} 


func (l *queue) deque()int{

	

	toremove := l.item[0]

	fmt.Println(toremove)

	l.item = l.item[1:]

	

	return toremove

}


func main(){

	do := queue{}
	fmt.Println(do)

	do.enque(10)
	do.enque(20)
	do.enque(30)
	do.enque(50)

	fmt.Println(do)

	do.deque()
	do.deque()
	fmt.Println(do)

}