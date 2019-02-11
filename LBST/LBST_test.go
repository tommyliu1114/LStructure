package LBST

import (
	"fmt"
	"testing"
)

type Lstring string

func (str Lstring) Compare(str1 interface{}) int {
	sstr := (string)(str)
	sstr1 := string(str1.(Lstring))
	if sstr < sstr1 {
		return -1
	}
	if sstr > sstr1 {
		return 1
	}
	return 0

}

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

func outNode(node *Node) {
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

	lbst1 := MkLBST()

	lbst1.Add(Lint(4))
	lbst1.Add(Lint(2))
	lbst1.Add(Lint(6))
	lbst1.Add(Lint(3))
	lbst1.Add(Lint(9))
	lbst1.Add(Lint(1))
	lbst1.Add(Lint(5))

	lbst1.Traverse(outNode)

	//fmt.Println(lbst1.Contains(Lint(2)))
	fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>")
	//lbst1.RemoveMin()
	lbst1.Remove(Lint(1))
	lbst1.Traverse(outNode)
	fmt.Println("the min of 4 2 6 3 9 1 5 is : ", lbst1.GetMin().Data, " now the size is : ", lbst1.GetSize())
	//lbst1.RemoveMax()
	lbst1.Remove(Lint(6))
	lbst1.Traverse(outNode)
	fmt.Println("the max of 4 2 6 3 9 1 5 is : ", lbst1.GetMax().Data, " now the size is : ", lbst1.GetSize())

}

func TestNoRecurTraverse(t *testing.T) {
	lbst1 := MkLBST()

	lbst1.Add(Lint(4))
	lbst1.Add(Lint(2))
	lbst1.Add(Lint(6))
	lbst1.Add(Lint(3))
	lbst1.Add(Lint(9))
	lbst1.Add(Lint(1))
	lbst1.Add(Lint(5))
	lbst1.NoRecurTraverse(outNode)
}

func TestBFS(t *testing.T) {
	lbst1 := MkLBST()

	lbst1.Add(Lint(4))
	lbst1.Add(Lint(2))
	lbst1.Add(Lint(6))
	lbst1.Add(Lint(3))
	lbst1.Add(Lint(9))
	lbst1.Add(Lint(1))
	lbst1.Add(Lint(5))
	fmt.Println("--------------")
	lbst1.BFS(outNode)
}
