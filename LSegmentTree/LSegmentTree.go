package LSegmentTree

import (
	LInterface "LStructure/LInterface"
)

type LSegmentTree struct {
	TreeData     []LInterface.LComparable
	OriginalData []LInterface.LComparable
	Size         int
	OptFunc      SegTreeOpt
}

type SegTreeOpt func(left LInterface.LComparable, right LInterface.LComparable) LInterface.LComparable

func MkLSegmentTree(data []LInterface.LComparable, optFunc SegTreeOpt) *LSegmentTree {
	ret := new(LSegmentTree)
	ret.OriginalData = data
	ret.Size = 4 * len(data)
	ret.OptFunc = optFunc
	ret.TreeData = make([]LInterface.LComparable, ret.Size)
	ret.buildSegmentTree(0, 0, ret.Size/4-1)
	return ret
}

func (seg *LSegmentTree) buildSegmentTree(treeIndex, start, end int) {
	//at treeIndex ,create segment originated from start to end
	if start == end {
		seg.TreeData[treeIndex] = seg.OriginalData[start]
		return
	}
	leftTreeIndex := seg.getLeftChild(treeIndex)
	rightTreeIndex := seg.getRightChild(treeIndex)
	mid := (start + end) / 2
	seg.buildSegmentTree(leftTreeIndex, start, mid)
	seg.buildSegmentTree(rightTreeIndex, mid+1, end)
	seg.TreeData[treeIndex] = seg.OptFunc(seg.TreeData[leftTreeIndex], seg.TreeData[rightTreeIndex])
}

func (seg *LSegmentTree) GetOriginalDataByIndex(index int) LInterface.LComparable {
	return seg.OriginalData[index]
}

func (seg *LSegmentTree) SetOriginalDataByIndex(index int, cdata LInterface.LComparable) {
	seg.OriginalData[index] = cdata
	seg.setData(0, 0, seg.Size/4-1, index, cdata)
}

func (seg *LSegmentTree) setData(treeindex, start, end, index int, cdata LInterface.LComparable) {
	if start == end {
		seg.TreeData[treeindex] = cdata
		return
	}
	mid := (start + end) / 2
	leftTreeIndex := seg.getLeftChild(treeindex)
	rightTreeIndex := seg.getRightChild(treeindex)
	if index >= mid+1 {
		seg.setData(rightTreeIndex, mid+1, end, index, cdata)
	} else if index <= mid {
		seg.setData(leftTreeIndex, start, mid, index, cdata)
	}
	seg.TreeData[treeindex] = seg.OptFunc(seg.TreeData[leftTreeIndex], seg.TreeData[rightTreeIndex])
}

func (seg *LSegmentTree) getLeftChild(index int) int {
	return 2*index + 1
}

func (seg *LSegmentTree) getRightChild(index int) int {
	return 2*index + 2
}

func (seg *LSegmentTree) GetSize() int {
	return seg.Size
}

func (seg *LSegmentTree) GetSeg(queryStart, queryEnd int) LInterface.LComparable {
	return seg.getSeg(0, 0, seg.Size/4-1, queryStart, queryEnd)
}
func (seg *LSegmentTree) getSeg(treeIndex, start, end, queryStart, queryEnd int) LInterface.LComparable {
	if start == queryStart && end == queryEnd {
		return seg.TreeData[treeIndex]
	}
	mid := (start + end) / 2
	leftTreeIndex := seg.getLeftChild(treeIndex)
	rightTreeIndex := seg.getRightChild(treeIndex)
	if queryStart >= mid+1 { //query right child
		return seg.getSeg(rightTreeIndex, mid+1, end, queryStart, queryEnd)
	} else if queryEnd <= mid { //query left child
		return seg.getSeg(leftTreeIndex, start, mid, queryStart, queryEnd)
	} else {
		leftReturn := seg.getSeg(leftTreeIndex, start, mid, queryStart, mid)
		rightReturn := seg.getSeg(rightTreeIndex, mid+1, end, mid+1, queryEnd)
		return seg.OptFunc(leftReturn, rightReturn)
	}
	return nil //just for grammer
}
