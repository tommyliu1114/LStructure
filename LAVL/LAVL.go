package LAVL

import (
	LInterface "LStructure/LInterface"
	LQue "LStructure/LQue"
)

type Node struct {
	Data   LInterface.LComparable //data
	Left   *Node
	Right  *Node
	Height int //sub-tree length,rooted at node

}

type LAVL struct {
	Root *Node
	Size int
}

type TraFunc func(*Node)

var checkArray []LInterface.LComparable

func MkLAVL() *LAVL {
	ret := new(LAVL)
	return ret
}

func (avl *LAVL) GetSize() int {
	return avl.Size
}

func (avl *LAVL) getHeight(node *Node) int {
	if node == nil {
		return 0
	}
	return node.Height
}

func (avl *LAVL) getBalanCoef(node *Node) int {
	if node == nil {
		return 0
	}
	return avl.getHeight(node.Left) - avl.getHeight(node.Right)
}

func (avl *LAVL) max(p, q int) int {
	ret := p
	if q > p {
		ret = q
	}
	return ret
}
func (avl *LAVL) min(p, q int) int {
	ret := p
	if q < p {
		ret = q
	}
	return ret
}

func (avl *LAVL) Remove(cdata LInterface.LComparable) {
	avl.remove(avl.Root, cdata)
}

func (avl *LAVL) remove(node *Node, cdata LInterface.LComparable) *Node {
	if node == nil {
		return nil
	}
	cp := cdata.Compare(node.Data)
	var retNode *Node
	if cp == -1 {
		node.Left = avl.remove(node.Left, cdata)
		retNode = node
	} else if cp == 1 {
		node.Right = avl.remove(node.Right, cdata)
		retNode = node
	} else if cp == 0 {
		if node.Left == nil {
			right := node.Right
			node.Right = nil
			avl.Size--
			retNode = right
		} else if node.Right == nil {
			left := node.Left
			node.Left = nil
			avl.Size--
			retNode = left
		} else if node.Right != nil && node.Left != nil {
			//here we can delte the left max or the right min,now we do the right min;
			lsuccess := new(Node)
			lsuccess.Data = avl.getMin(node.Right).Data

			lsuccess.Left = node.Left
			lsuccess.Right = avl.remove(node.Right, lsuccess.Data)

			node.Right = nil
			node.Left = nil

			retNode = lsuccess
		}

	}
	if retNode == nil {
		return nil
	}

	retNode.Height = 1 + avl.max(avl.getHeight(retNode.Left), avl.getHeight(retNode.Right))
	balanCoef := avl.getBalanCoef(retNode)

	if balanCoef > 1 && avl.getBalanCoef(retNode.Left) >= 0 {
		//left tree too deep,because  left's left-child is added(LL) right turn(LL)
		retNode = avl.rightRotate(retNode)
	} else if balanCoef < -1 && avl.getBalanCoef(retNode.Right) <= 0 {
		//right tree too deep,because  right's right-child is added(RR) right turn(RR)
		retNode = avl.leftRotate(retNode)

	} else if balanCoef > 1 && avl.getBalanCoef(retNode.Left) < 0 {
		//left tree too deep ,because  left's right-child is added(LR)
		retNode.Left = avl.leftRotate(retNode.Left)
		retNode = avl.rightRotate(retNode)

	} else if balanCoef < -1 && avl.getBalanCoef(retNode.Right) > 0 {
		//left tree too deep ,because  right's left-child is added(RL)
		retNode.Right = avl.rightRotate(retNode.Right)
		retNode = avl.leftRotate(retNode)
	}

	return retNode
}
func (avl *LAVL) Add(cdata LInterface.LComparable) {
	avl.Root = avl.add(avl.Root, cdata)
}

func (avl *LAVL) add(node *Node, cdata LInterface.LComparable) *Node {
	if node == nil {

		avl.Size++
		ret := new(Node)
		ret.Data = cdata
		ret.Height = 1 //initail leaf's height is 1;

		return ret
	}
	cp := cdata.Compare(node.Data)
	if cp == -1 {
		node.Left = avl.add(node.Left, cdata)
	} else if cp == 1 {
		node.Right = avl.add(node.Right, cdata)
	}
	node.Height = 1 + avl.max(avl.getHeight(node.Left), avl.getHeight(node.Right))
	balanCoef := avl.getBalanCoef(node)

	if balanCoef > 1 && avl.getBalanCoef(node.Left) >= 0 {
		//left tree too deep,because  left's left-child is added(LL) right turn(LL)
		node = avl.rightRotate(node)
	} else if balanCoef < -1 && avl.getBalanCoef(node.Right) <= 0 {
		//right tree too deep,because  right's right-child is added(RR) right turn(RR)
		node = avl.leftRotate(node)

	} else if balanCoef > 1 && avl.getBalanCoef(node.Left) < 0 {
		//left tree too deep ,because  left's right-child is added(LR)
		node.Left = avl.leftRotate(node.Left)
		node = avl.rightRotate(node)

	} else if balanCoef < -1 && avl.getBalanCoef(node.Right) > 0 {
		//left tree too deep ,because  right's left-child is added(RL)
		node.Right = avl.rightRotate(node.Right)
		node = avl.leftRotate(node)
	}
	return node
}

func (avl *LAVL) leftRotate(node *Node) *Node {
	newRoot := node.Right
	node.Right = newRoot.Left
	newRoot.Left = node
	//update height ,attention the update order
	node.Height = avl.max(avl.getHeight(node.Left), avl.getHeight(node.Right)) + 1
	newRoot.Height = avl.max(avl.getHeight(newRoot.Left), avl.getHeight(newRoot.Right)) + 1
	return newRoot
}

func (avl *LAVL) rightRotate(node *Node) *Node {

	newRoot := node.Left
	node.Left = newRoot.Right
	newRoot.Right = node
	//update height,attention the update order
	node.Height = avl.max(avl.getHeight(node.Left), avl.getHeight(node.Right)) + 1
	newRoot.Height = avl.max(avl.getHeight(newRoot.Left), avl.getHeight(newRoot.Right)) + 1
	return newRoot
}

func (avl *LAVL) IsBST() bool {
	checkArray = make([]LInterface.LComparable, 0, avl.Size) //everytime clear the slice
	//fmt.Println("avl size is : ", avl.Size)
	ret := true
	avl.Traverse(func(node *Node) {
		checkArray = append(checkArray, node.Data)
		return
	}, "inorder")
	//fmt.Println(checkArray)
	for i := 1; i < len(checkArray); i++ {
		cp := checkArray[i-1].Compare(checkArray[i])
		if cp > 1 {
			ret = false
			break
		}
	}
	return ret
}

func (avl *LAVL) IsBalanced() bool {
	return avl.isBalanced(avl.Root)
}
func (avl *LAVL) isBalanced(node *Node) bool {
	if node == nil {
		return true
	}
	if avl.getBalanCoef(node) > 1 {
		return false
	}
	return avl.isBalanced(node.Left) && avl.isBalanced(node.Right)
}

func (avl *LAVL) Traverse(tra TraFunc, traverseType string) {

	if "inorder" == traverseType {
		avl.inorderTraverse(avl.Root, tra)
		return
	}
	avl.traverse(avl.Root, tra)

}

func (avl *LAVL) inorderTraverse(node *Node, tra TraFunc) {
	if node == nil {
		return
	}
	//tra(node) //preorder
	avl.traverse(node.Left, tra)
	tra(node) //inorder
	avl.traverse(node.Right, tra)
	//tra(node) //postorder
}
func (avl *LAVL) traverse(node *Node, tra TraFunc) {
	if node == nil {
		return
	}
	//tra(node) //preorder
	avl.traverse(node.Left, tra)
	tra(node) //inorder
	avl.traverse(node.Right, tra)
	//tra(node) //postorder
}

func (avl *LAVL) BFS(tra TraFunc) {
	//broad first
	que := LQue.MakeQue()
	if avl.Root == nil {
		return
	}
	que.Enque(avl.Root)
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

func (avl *LAVL) Contains(cdata LInterface.LComparable) bool {
	return avl.constains(avl.Root, cdata)
}

func (avl *LAVL) constains(node *Node, cdata LInterface.LComparable) bool {
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

		isExist = avl.constains(node.Left, cdata)
	} else if cp == 1 {

		isExist = avl.constains(node.Right, cdata)
	}
	return isExist
}

func (avl *LAVL) GetMin() *Node {
	return avl.getMin(avl.Root)
}
func (avl *LAVL) getMin(node *Node) *Node {
	//find the minimum of the tree whose root is the ndoe
	if node == nil {
		return nil
	}
	var imin *Node = node
	if node.Left != nil {
		lmin := avl.getMin(node.Left)

		if lmin.Data.Compare(imin.Data) == -1 {
			imin = lmin
		}
	}
	return imin

}

func (avl *LAVL) GetMax() *Node {
	return avl.getMax(avl.Root)
}
func (avl *LAVL) getMax(node *Node) *Node {
	if node == nil {
		return nil
	}
	var imax *Node = node
	if node.Right != nil {
		rmax := avl.getMax(node.Right)
		if rmax.Data.Compare(imax.Data) == 1 {
			imax = rmax
		}
	}
	return imax
}
