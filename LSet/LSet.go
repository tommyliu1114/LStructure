package LSet

import (
	LBST "LStructure/LBST"
	LInterface "LStructure/LInterface"
)

/*
here use BST as the foundation ,later can alter the foundation into AVL similarly
*/
type LSet LBST.LBST

func MkSet() *LSet {
	set := (*LSet)(LBST.MkLBST())
	return set
}
func (lset *LSet) GetSize() int {
	return (*LBST.LBST)(lset).GetSize()
}

func (lset *LSet) Contains(cdata LInterface.LComparable) bool {
	return ((*LBST.LBST)(lset)).Contains(cdata)
}

func (lset *LSet) Remove(cdata LInterface.LComparable) {
	(*LBST.LBST)(lset).Remove(cdata)
	return
}

func (lset *LSet) Add(cdata LInterface.LComparable) {
	(*LBST.LBST)(lset).Add(cdata)
	return
}
