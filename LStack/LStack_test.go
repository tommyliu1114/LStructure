package LStack

import (
	List "LStructure/LLinkedList"
	"fmt"
	"testing"
)

func TestPush(t *testing.T) {
	stack := MakeStack()
	stack.Push("Happy")
	stack.Push("NewYear")
	//(*List.LLinkedList)(stack).Trans(trans)
}

func TestPop(t *testing.T) {
	stack := MakeStack()
	stack.Push("Happy")
	stack.Push("NewYear")
	stack.Push("LewisTier")
	(*List.LLinkedList)(stack).Trans(trans)
	edata := stack.Pop()
	fmt.Print("\nHere we get poped data : ")
	fmt.Println(edata)
	fmt.Println("=========================")
	//(*List.LLinkedList)(stack).Trans(trans)
	fmt.Println("===========peek the top gun================")
	pdata := stack.Peek()
	fmt.Println(pdata)
	stack.Push("the next one top ")
	fmt.Println("after push one elements")
	(*List.LLinkedList)(stack).Trans(trans)
	fmt.Println("\nall stack's size is : ", stack.GetSize())

}

func trans(cnode *List.Node) {
	fmt.Print(" --->")
	fmt.Print(cnode.Data)

}
