package LBST

import (
	LInterface "LStructure/LInterface"

	LStack "LStructure/LStack"

	LQue "LStructure/LQue"
)

type Node struct {
	Data  LInterface.LComparable //data
	Left  *Node
	Right *Node
}

type LBST struct {
	Root *Node
	Size int
}

func MkLBST() *LBST {
	bst := new(LBST)
	return bst
}

type TraFunc func(*Node)

func (lbst *LBST) GetSize() int {
	return lbst.Size
}

func (lbst *LBST) Add(cdata LInterface.LComparable) {
	lbst.Root = lbst.add(lbst.Root, cdata)
}

func (lbst *LBST) add(node *Node, cdata LInterface.LComparable) *Node {
	if node == nil {
		lbst.Size++
		ret := new(Node)
		ret.Data = cdata
		return ret
	}
	cp := cdata.Compare(node.Data)
	if cp == -1 {
		node.Left = lbst.add(node.Left, cdata)
	} else if cp == 1 {
		node.Right = lbst.add(node.Right, cdata)
	}
	return node
}

func (lbst *LBST) Contains(cdata LInterface.LComparable) bool {
	return lbst.constains(lbst.Root, cdata)
}

func (lbst *LBST) constains(node *Node, cdata LInterface.LComparable) bool {
	var isExist bool
	var cp int
	if node != nil {
		cp = cdata.Compare(node.Data)

	}
	if node == nil {
		isExist = false
	} else if cp == 0 {
		isExist = true

	} else if cp == -1 {

		isExist = lbst.constains(node.Left, cdata)
	} else if cp == 1 {

		isExist = lbst.constains(node.Right, cdata)
	}
	return isExist
}

func (lbst *LBST) Traverse(tra TraFunc) {
	lbst.traverse(lbst.Root, tra)

}
func (lbst *LBST) traverse(node *Node, tra TraFunc) {
	if node == nil {
		return
	}
	tra(node) //preorder
	lbst.traverse(node.Left, tra)
	//tra(node) //inorder
	lbst.traverse(node.Right, tra)
	//tra(node) //postorder
}

func (lbst *LBST) NoRecurTraverse(tra TraFunc) {
	stack := LStack.MakeStack()
	if lbst.Root == nil {
		return
	}

	stack.Push(lbst.Root)
	for !stack.IsEmpty() {
		cur := stack.Pop()
		tra(cur.(*Node))
		if (cur.(*Node)).Right != nil {
			stack.Push(cur.(*Node).Right)
		}
		if (cur.(*Node)).Left != nil {
			stack.Push((cur.(*Node)).Left)
		}
	}

}

func (lbst *LBST) BFS(tra TraFunc) {
	//broad first
	que := LQue.MakeQue()
	if lbst.Root == nil {
		return
	}
	que.Enque(lbst.Root)
	for !que.IsEmpty() {
		cur, _ := que.Deque()
		tra(cur.(*Node))
		if (cur.(*Node).Left) != nil {
			que.Enque(cur.(*Node).Left)
		}
		if (cur.(*Node).Right) != nil {
			que.Enque(cur.(*Node).Right)
		}
	}
}

func (lbst *LBST) GetMin() *Node {
	return lbst.getMin(lbst.Root)
}
func (lbst *LBST) getMin(node *Node) *Node {
	//find the minimum of the tree whose root is the ndoe
	if node == nil {
		return nil
	}
	var imin *Node = node
	if node.Left != nil {
		lmin := lbst.getMin(node.Left)

		if lmin.Data.Compare(imin.Data) == -1 {
			imin = lmin
		}
	}
	return imin

}

func (lbst *LBST) GetMax() *Node {
	return lbst.getMax(lbst.Root)
}
func (lbst *LBST) getMax(node *Node) *Node {
	if node == nil {
		return nil
	}
	var imax *Node = node
	if node.Right != nil {
		rmax := lbst.getMax(node.Right)
		if rmax.Data.Compare(imax.Data) == 1 {
			imax = rmax
		}
	}
	return imax
}
func (lbst *LBST) Remove(cdata LInterface.LComparable) {
	lbst.Root = lbst.remove(lbst.Root, cdata)
}

func (lbst *LBST) remove(node *Node, cdata LInterface.LComparable) *Node {
	if node == nil {
		return nil
	}
	cp := cdata.Compare(node.Data)
	if cp == -1 {
		node.Left = lbst.remove(node.Left, cdata)
		return node
	} else if cp == 1 {
		node.Right = lbst.remove(node.Right, cdata)
		return node
	} else if cp == 0 {
		if node.Left == nil {
			right := node.Right
			node.Right = nil
			lbst.Size--
			return right
		}
		if node.Right == nil {
			left := node.Left
			node.Left = nil
			lbst.Size--
			return left
		}
		//here we can delte the left max or the right min,now we do the left max;
		lsuccess := lbst.getMax(node.Left)
		lsuccess.Left = lbst.removeMax(node.Left)
		lsuccess.Right = node.Right
		node.Right = nil
		node.Left = nil
		return lsuccess

	}
	return nil //for grammer
}

func (lbst *LBST) RemoveMax() {
	lbst.removeMax(lbst.Root)
}
func (lbst *LBST) removeMax(node *Node) *Node {
	//remove the max of the tree whose root is node
	if node.Right == nil {
		left := node.Left
		node.Left = nil
		lbst.Size--
		return left
	}
	node.Right = lbst.removeMax(node.Right)
	return node
}

func (lbst *LBST) removeMin(node *Node) *Node {
	if node.Left == nil {
		right := node.Right
		node.Right = nil
		lbst.Size--
		return right
	}
	node.Left = lbst.removeMin(node.Left)
	return node
}

func (lbst *LBST) RemoveMin() {
	lbst.removeMin(lbst.Root)
}
