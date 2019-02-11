package LHeap

import (
	LInterface "LStructure/LInterface"
)

type LHeap []LInterface.LComparable

func MkHeap() *LHeap {

	var ret LHeap
	return &ret
}

func (heap *LHeap) GetSize() int {
	length := len(([]LInterface.LComparable)(*heap))
	return length
}

func (heap *LHeap) getDataByIndex(index int) LInterface.LComparable {
	ret := ([]LInterface.LComparable(*heap))[index]
	return ret
}

func (heap *LHeap) setDataByIndex(index int, cdata LInterface.LComparable) {
	[]LInterface.LComparable(*heap)[index] = cdata
	return
}

func (heap *LHeap) getParent(index int) int {
	//heap begin from 0
	return (index - 1) / 2
}

func (heap *LHeap) getLeftChild(index int) int {
	return index*2 + 1
}
func (heap *LHeap) getRightChild(index int) int {
	return index*2 + 2
}

func (heap *LHeap) Add(cdata LInterface.LComparable) {
	/* add one element ,two stpes :
	1,add one element at last
	2,adjust the element
	*/
	*heap = append(*heap, cdata)
	heap.adjustUp(heap.GetSize() - 1)
}

func (heap *LHeap) Peek() LInterface.LComparable {
	return heap.getDataByIndex(0)
}

func (heap *LHeap) Remove() LInterface.LComparable {
	/*take the max(min) element, two stpes:
	1,take the element 0
	2,exchange element 0 with element size-1
	3,squeeze the heap from 0 to size-1
	2,adjust the element 0
	*/
	ret := heap.getDataByIndex(0)
	heap.exchange(0, heap.GetSize()-1)
	*heap = (LHeap)(([]LInterface.LComparable)(*heap)[0 : heap.GetSize()-1])
	heap.adjustDown(0)
	return ret

}
func (heap *LHeap) exchange(i int, j int) {
	tem := heap.getDataByIndex(i)
	heap.setDataByIndex(i, heap.getDataByIndex(j))
	heap.setDataByIndex(j, tem)
	return
}

func (heap *LHeap) adjustUp(index int) {
	for index > 0 && heap.getDataByIndex(index).Compare(heap.getDataByIndex(heap.getParent(index))) == 1 {
		heap.exchange(index, heap.getParent(index))
		index = heap.getParent(index)
	}
}

func (heap *LHeap) adjustDown(index int) {
	for heap.getLeftChild(index) < heap.GetSize() {
		maxInd := heap.getLeftChild(index)
		if heap.getRightChild(index) < heap.GetSize() {
			//if right element exist
			if heap.getDataByIndex(heap.getRightChild(index)).Compare(heap.getDataByIndex(heap.getLeftChild(index))) == 1 {
				maxInd = heap.getRightChild(index)
			}

		}
		if heap.getDataByIndex(index).Compare(heap.getDataByIndex(maxInd)) == 1 {
			break //if index element is greater than the max(rightchild,leftchild),terminate
		}
		heap.exchange(index, maxInd)
		index = maxInd
	}
	return
}
