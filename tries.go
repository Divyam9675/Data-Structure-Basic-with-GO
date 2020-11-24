package main

import "fmt"


const Size = 23

type node struct{

 childeren [23]*node
 isEnd  bool

}

type tries struct{

	root *node

}


func relation()*tries{

	res := &tries{root: &node{}}

	return res
}


func (l *tries)insert(n string){

        total  := len(n)

		current := l.root
		//fmt.Println(current)

		for i:=0; i<total; i++{

			temp := n[i] - 'a'

			if current.childeren[temp] == nil{

				current.childeren[temp] = &node{}

			}

			current = current.childeren[temp]
		
		}

		current.isEnd = true
}

func main(){

	//pos := &tries{}

	pos := relation()

	//fmt.Println(pos)

	arr  := []string{
		"aragorn",
		"aragon",
		"argon",
		"eragon",
		"oregon",
		"oregano",
		"oreo",
	}

	for _ , val := range arr{

		fmt.Println(val)

	    pos.insert(val)	

	} 
}