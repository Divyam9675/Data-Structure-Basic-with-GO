package main

import "fmt"


// 1) Describe the node first

type node struct{

	data int
	next *node

	// next contain address of the next node that's why we use pointer to the node
}


// 2. Now we need to describe list 

type list struct{

	head *node 

	// head indicates the starting node of the list 

	length int


	// length indicate the size of the list

}


// 3. Now we put node in to the list

func (l *list)insert(n *node){

	// 1. Firstly put create a temporary space and on that space we put header

	second := l.head
	
	// 2. Now set the new node as head

	l.head = n

	// 3. let  the point old head as a next node of head

	l.head.next = second

	l.length++


}


//4. Print the data of whole list in a sequence

func (l list)print(){

	// To print full node we use for loop to iterate 

	// And count the list untill list become 0


	toprint := l.head

	for l.length != 0{

		// Inside for we need to print first node But
		// If we just print head we get only head not full list we need to follow up for complete list thats why
		// we declare a variable outside for loop where we store head 

	//	fmt.Printf("%d", l.head.data)

		fmt.Printf("%d", toprint.data)

		// After print we need to set node updated


		toprint = toprint.next

	//	l.head = l.head.next

		// Now we decrease the node

		l.length--

	}
 

	fmt.Println("\n")

}

// 5) TO delete the given data from the list

func (l *list)delete(value int){


	// 6 . Return when search value for delete is on the head

	if l.head.data == value {

		l.head  = l.head.next

		l.length--

		return
	}


	// 7. if the list empty and we searching values on node

	if l.length == 0{

		return

	}


	// we need to create a variable to store the head value

	todelete := l.head

	// SO here we dont want to compare the data of the node with todelete variable , we need to compare the data with the next of the node

	for todelete.next.data != value{


		// 8. If searching element not found on the list

		if todelete.next.next == nil{
			return
		}


		// IF data is not match we just with next node

		todelete = todelete.next


	} 

	// if data found skip that node from the list and updata the node

	todelete.next = todelete.next.next
	l.length--

}












func main(){

	mylist := list{}
	mynode := &node{data: 400}
	mynode1 := &node{data: 500}
	mynode2 := &node{data: 300}
	mynode3 := &node{data: 200}
	mynode4 := &node{data: 100}
	mynode5 := &node{data: 800}
	mynode6 := &node{data: 700}

	mylist.insert(mynode)
	mylist.insert(mynode1)
	mylist.insert(mynode2)
	mylist.insert(mynode3)
	mylist.insert(mynode4)
	mylist.insert(mynode5)
	mylist.insert(mynode6)

	//fmt.Println(mylist)

	mylist.print()
	mylist.delete(300)
	mylist.print()

	mylist.delete(2)
	mylist.print()

	//nod := list{}

	//nod.delete(20)

	//nod.print()

	mylist.delete(700)
	//mylist.print()


}






