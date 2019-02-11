package LAVL

import (
	"fmt"
	"math/rand"
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

func outNode(node *Node) {
	//fmt.Println("=====begin=====")
	fmt.Printf("address : %p ", node)
	fmt.Println(" and data : ", *node)
	//fmt.Println("======end=========")
}
func TestIsBST(t *testing.T) {
	avl := MkLAVL()
	for i := 0; i < 20; i++ {
		avl.Add(Lint(rand.Int() % 20))
	}
	avl.BFS(outNode)
	fmt.Println("=================")

	avl.Traverse(outNode, "inorder")
	fmt.Println("++++++++++++++++++")
	for i := 0; i < 5; i++ {
		avl.Remove(Lint(rand.Int() % 20))
	}
	avl.BFS(outNode)
	for i := 0; i < 10; i++ {
		k := rand.Int() % 20
		fmt.Println(k, " is contained", avl.Contains(Lint(k)))
	}
	/*
		avl.Add(Lint(4))
		avl.Add(Lint(2))
		avl.Add(Lint(6))
		avl.Add(Lint(3))
		avl.Add(Lint(9))
		avl.Add(Lint(1))
		avl.Add(Lint(5))
		avl.Add(Lint(10))
		avl.Add(Lint(11))
		avl.Add(Lint(12))
		avl.Add(Lint(13))
		avl.Add(Lint(7))
		avl.Add(Lint(8))
		avl.Add(Lint(-8))
		avl.Add(Lint(-2))
		fmt.Println("avl is bst ", avl.IsBST())
		fmt.Println("avl is balanced : ", avl.IsBalanced())
		fmt.Println("avl is : ")
		avl.BFS(outNode)
		fmt.Println("++++++++++++++++++++++++")
		avl.Remove(Lint(8))
		avl.Remove(Lint(2))

		avl.Remove(Lint(10))
		fmt.Println("avl is bst ", avl.IsBST())
		fmt.Println("avl is balanced : ", avl.IsBalanced())
		fmt.Println("delete 2 ,8 ,10 is ")
		avl.BFS(outNode)
		fmt.Println("=================")

		avl.Traverse(outNode, "inorder")
	*/
}
