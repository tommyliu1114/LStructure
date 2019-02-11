package LLinkedList

import(
	"errors"
)

type Node struct{
	Data interface{}
	next *Node
}
type LLinkedList struct{
	DummyHead *Node
	//tail *Node
	size  int
	//no dummy nodes
}
func (list *LLinkedList) GetSize() int{
	return list.size
}
func MakeLink() *LLinkedList{
	link := new(LLinkedList)
	link.size = 0
	link.DummyHead = new(Node)
	return link
}

func (list *LLinkedList) Trans(trans func(node *Node)){
	cur := list.DummyHead.next
	for cur != nil {
		trans(cur)
		cur = cur.next
	} 
}

func (list *LLinkedList) IsEmpty() bool{
	return (list.size == 0)
}

func (list *LLinkedList) InsertFirst(cdata interface{}){
	list.InsertByIndex(cdata,0)
	
}

func (list *LLinkedList) InsertLast(cdata interface{}){
	list.InsertByIndex(cdata,list.size)
}

func (list *LLinkedList) InsertByIndex(cdata interface{},ind int) error{
	if ind < 0 || ind > list.size{
		return errors.New("index error")
	}
	cur := list.DummyHead
	var i int
	for i = 0 ; i < ind ; i++{
		cur = cur.next
	}
	node := new(Node)
	node.Data = cdata
	node.next = cur.next
	cur.next = node
	list.size++
	return nil
	
}

func (list *LLinkedList) GetByIndex(ind int) (node *Node,err error){
	if ind < 0 || ind >= list.size{
		err = errors.New("index error")
		node = nil
		return
	}
	cur := list.DummyHead.next
	var i int
	for i = 0; i< ind;i++{
		cur = cur.next
	}
	node = cur
	return 
}

func (list *LLinkedList) SetByIndex(cdata interface{},ind int) error{
	if ind < 0 || ind >= list.size{
		err := errors.New("index error")
		return err
	}
	cur := list.DummyHead.next
	var i int
	for i = 0; i<ind;i++{
		cur = cur.next
	}
	cur.Data = cdata
	return nil
}

func (list *LLinkedList) Contains(cdata interface{}) bool{
	cur := list.DummyHead.next
	for ;cur != nil ;cur=cur.next{
		if cur.Data == cdata{
			return true
		}
	}
	return false
}
func (list *LLinkedList) DeleteByIndex(ind int) (node *Node,err error){
	if ind < 0 || ind >= list.size{
		err = errors.New("index error")
		node = nil
		return
	}
	var i int
	prev := list.DummyHead
	for i=0;i<ind;i++{
		prev = prev.next
	}
	node = prev.next
	prev.next = node.next
	list.size--
	return node,nil
}

func (list *LLinkedList) DeleteFirst() (node *Node,err error){
	return list.DeleteByIndex(0)
}

func (list *LLinkedList) DeleteLast()(node *Node,err error){
	return list.DeleteByIndex(list.size-1)
}







