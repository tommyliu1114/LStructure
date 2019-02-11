package LPriorityQue

import (
	LHeap "LStructure/LHeap"
	LInterface "LStructure/LInterface"
)

type LPriorityQue LHeap.LHeap

func MkPriorityQue() *LPriorityQue {
	ret := LHeap.MkHeap()
	return (*LPriorityQue)(ret)
}

func (que *LPriorityQue) GetSize() int {
	return (*LHeap.LHeap)(que).GetSize()
}
func (que *LPriorityQue) Enque(cdata LInterface.LComparable) {
	(*LHeap.LHeap)(que).Add(cdata)
}
func (que *LPriorityQue) Deque() LInterface.LComparable {
	return (*LHeap.LHeap)(que).Remove()
}

func (que *LPriorityQue) GetFront() LInterface.LComparable {
	return (*LHeap.LHeap)(que).Peek()
}
