package LStack

import (
	List "LStructure/LLinkedList"
)


type LStack List.LLinkedList
/*
func init(){
	fmt.Println("here has loaded package")
}
*/

func MakeStack() *LStack{
	stack := List.MakeLink()
	return ((*LStack) (stack))
}

func (stack *LStack)IsEmpty() bool{
	return ((*List.LLinkedList)(stack)).IsEmpty()
}

func (stack *LStack)Push(cdata interface{}){
	((*List.LLinkedList)(stack)).InsertLast(cdata)
}

func (stack *LStack)Pop() interface{}{
	node,_ := ((*List.LLinkedList)(stack)).DeleteLast()
	return node.Data
}

func (stack *LStack)Peek() interface{}{
	clist := (*List.LLinkedList)(stack)
	node , _ := clist.GetByIndex(clist.GetSize()-1)
	return node.Data
}

func (stack *LStack)GetSize() int{
	clist := (*List.LLinkedList)(stack)
	return (clist.GetSize())
}
