package LQue

import (
	List "LStructure/LLinkedList"
	"fmt"
	"testing"
)

func trans(cnode *List.Node) {
	fmt.Print(" --->")
	fmt.Print(cnode.Data)

}

func TestEnque(t *testing.T) {
	var i int
	que := MakeQue()
	for i = 0; i < 10; i++ {
		que.Enque(i + 100)
	}
	((*List.LLinkedList)(que)).Trans(trans)
	fmt.Println("========================")

	for i = 0; i < 3; i++ {
		iq, _ := que.Deque()
		fmt.Println("the  ", i, " value is : ", iq)

	}
	((*List.LLinkedList)(que)).Trans(trans)
}
