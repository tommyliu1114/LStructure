package LLinkedList

import(
	"fmt"
	"testing"
)
func TestTrans(t *testing.T){

	var list *LLinkedList = MakeLink()
	istr1 := "Happy"
	 
	istr2 := "NewYear"
	list.InsertFirst(istr1)
	list.InsertFirst(istr2)
	list.Trans(trans)



	var list2 *LLinkedList = MakeLink()
	idata1 := 1
	idata2 := 2
	list2.InsertLast(idata1)
	list2.InsertLast(idata2)
	list2.Trans(trans)



	
}

func TestInsertByhIndex(t *testing.T){
	var list *LLinkedList = MakeLink()
	idata1 := 100.3
	idata2 := 200.3
	list.InsertLast(idata1)
	list.InsertLast(idata2)
	list.InsertByIndex(150.3,1)
	list.InsertByIndex(300.9,3)
	list.Trans(trans)
	list.DeleteByIndex(2)
	fmt.Println("=============after delete=======")
	list.Trans(trans)
}

func TestGetByIndex(t *testing.T){
	var list *LLinkedList = MakeLink()
	idata1 := 2.3
	idata2 := 5.3
	list.InsertLast(idata1)
	list.InsertLast(idata2)
	//list.Trans(trans)
	node1,_ := list.GetByIndex(0)
	node2,_ := list.GetByIndex(1)
	fmt.Println(node1.Data)
	fmt.Println(node2.Data)
}
func TestInsertLast(t *testing.T){
	var list *LLinkedList = MakeLink()
	idata1 := 1
	idata2 := 2
	list.InsertLast(idata1)
	list.InsertLast(idata2)
	list.Trans(trans)
}
func trans(cnode *Node){
	fmt.Print(" ")
	fmt.Println(cnode.Data)
}
func TestInsertFirst(t *testing.T){
	var list *LLinkedList = MakeLink()
	istr1 := "Happy"
	 
	istr2 := "NewYear"
	list.InsertFirst(istr1)
	list.InsertFirst(istr2)
	list.Trans(trans)
}