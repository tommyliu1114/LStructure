package LHeap

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func init() {
	//以时间作为初始化种子
	rand.Seed(time.Now().UnixNano())
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

func TestRemove(t *testing.T) {

	heap := MkHeap()
	for i := 0; i < 10; i++ {
		heap.Add(Lint(rand.Int() % 100))

	}
	fmt.Println(heap)

	for i := 0; i < 10; i++ {

		fmt.Println("now the value is : ", heap.Remove())
	}

}
