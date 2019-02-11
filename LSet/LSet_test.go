package LSet

import (
	LBST "LStructure/LBST"
	"fmt"
	"testing"
)

type Lint int

func (num Lint) Compare(num1 interface{}) int {
	nnum := (int)(num)

	nnum1 := int(num1.(Lint))
	//
	//fmt.Println("===", nnum, nnum1)
	if nnum < nnum1 {
		return -1
	} else if nnum == nnum1 {
		return 0
	} else {
		return 1
	}
}

func outNode(node *LBST.Node) {
	//fmt.Println("=====begin=====")
	fmt.Printf("address : %p ", node)
	fmt.Println(" and data : ", *node)
	//fmt.Println("======end=========")
}

func TestAdd(t *testing.T) {
	/*
		fmt.Println("================")
		lbst := MkLBST()
		str1 := (Lstring)("lewis")
		str2 := (Lstring)("letiger")
		str3 := (Lstring)("sen")
		lbst.Add(str1)
		lbst.Add(str2)
		lbst.Add(str3)
		lbst.Traverse(outNode)
		fmt.Println(lbst.Contains((Lstring)("lewis")))
		fmt.Println(lbst.Contains((Lstring)("tiger")))
		fmt.Println("+++++++++++++++++++++++++++++++++")
	*/

	lbst1 := MkSet()

	lbst1.Add(Lint(4))
	lbst1.Add(Lint(2))
	lbst1.Add(Lint(6))
	lbst1.Add(Lint(3))
	lbst1.Add(Lint(9))
	lbst1.Add(Lint(1))
	lbst1.Add(Lint(5))

	((*LBST.LBST)(lbst1)).Traverse(outNode)

	//fmt.Println(lbst1.Contains(Lint(2)))
	fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>")
	//lbst1.RemoveMin()
	lbst1.Remove(Lint(1))
	((*LBST.LBST)(lbst1)).Traverse(outNode)

	//lbst1.RemoveMax()
	lbst1.Remove(Lint(6))
	((*LBST.LBST)(lbst1)).Traverse(outNode)

}
