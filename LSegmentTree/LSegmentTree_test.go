package LSegmentTree

import (
	LInterface "LStructure/LInterface"
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

func sumOpt(left, right LInterface.LComparable) LInterface.LComparable {
	left1, _ := left.(Lint)
	left2, _ := right.(Lint)
	return left1 + left2
}

func maxOpt(left, right LInterface.LComparable) LInterface.LComparable {
	left1, _ := left.(Lint)
	right1, _ := right.(Lint)
	var ret Lint = left1
	if right1 > left1 {
		ret = right1
	}
	return ret

}

func TestMkLSegmentTree(t *testing.T) {
	a := []LInterface.LComparable{Lint(-2), Lint(0), Lint(3), Lint(-5), Lint(2), Lint(-1)}

	seg := MkLSegmentTree(a, SegTreeOpt(sumOpt))
	fmt.Println("original data is : ", seg.OriginalData)
	fmt.Println("tree data is : ", seg.TreeData)
	fmt.Println("get seq of 0,3 is : ", seg.GetSeg(0, 3))
	fmt.Println("=====================")
	seg1 := MkLSegmentTree(a, SegTreeOpt(maxOpt))
	fmt.Println("original data is : ", seg1.OriginalData)
	fmt.Println("tree data is : ", seg1.TreeData)
	fmt.Println("get seq of 3,5 is : ", seg1.GetSeg(3, 5))
	fmt.Println("+++++++++++++++++++++++++++++")
	seg2 := MkLSegmentTree(a, SegTreeOpt(maxOpt))
	fmt.Println("original data is : ", seg2.OriginalData)
	fmt.Println("tree data is : ", seg2.TreeData)
	seg2.SetOriginalDataByIndex(3, Lint(-7))
	fmt.Println("get seq of 3,5 is : ", seg2.GetSeg(3, 5))
}
