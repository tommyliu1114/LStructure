package LQue

import (
	List "LStructure/LLinkedList"
)

type LQue List.LLinkedList

func MakeQue() *LQue {
	que := List.MakeLink()
	return ((*LQue)(que))
}

func (que *LQue) IsEmpty() bool {
	cl := (*List.LLinkedList)(que)
	return cl.IsEmpty()

}

func (que *LQue) Enque(cdata interface{}) {
	cl := (*List.LLinkedList)(que)
	cl.InsertLast(cdata)
}

func (que *LQue) Deque() (cdata interface{}, err error) {
	cl := (*List.LLinkedList)(que)
	cnode, cerr := cl.DeleteFirst()
	if cerr != nil {
		err = cerr
		cdata = nil
		return
	}
	cdata = cnode.Data
	return

}

func (que *LQue) GetSize() int {
	cl := (*List.LLinkedList)(que)
	return cl.GetSize()
}
