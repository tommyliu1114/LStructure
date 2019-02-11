package LUnionFind

import (
	"fmt"
	"testing"
)

func TestUnion(t *testing.T) {
	fmt.Println("=======================")
	uf := MkUnionFind(20)
	uf.Union(3, 9)
	uf.Union(2, 9)
	uf.Union(4, 2)
	fmt.Println("3 and 9 connected is ", uf.IsConnnected(3, 9))
	fmt.Println("3 and 4 connected is :", uf.IsConnnected(3, 4))
	fmt.Println(uf.ParentId)
	fmt.Println(uf.TreeWeight)

}
func TestIsConnnected(t *testing.T) {

}
