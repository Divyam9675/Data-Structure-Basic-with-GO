package main

import "fmt"

type stack struct{

	item[] int

}

func (s *stack) Push(i int){

	s.item = append(s.item, i)

}

func (s *stack) Pop()int{


	l := len(s.item)-2

    todelete := s.item[l]

	s.item	= s.item[:l]
	
	


  return todelete

}

func main(){


	on := stack{}

	fmt.Println(on)

	on.Push(10)
	on.Push(20)
	on.Push(23)
	on.Push(44)
	on.Push(46)

	fmt.Println(on)

	on.Pop()

	fmt.Println(on)
}