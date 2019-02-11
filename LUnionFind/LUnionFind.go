package LUnionFind

type LUnionFind struct {
	ParentId   []int
	Size       int
	TreeWeight []int
}

func MkUnionFind(size int) *LUnionFind {
	ret := new(LUnionFind)
	ret.Size = size
	ret.ParentId = make([]int, ret.Size)
	ret.TreeWeight = make([]int, ret.Size)
	for i := 0; i < ret.Size; i++ {
		ret.ParentId[i] = i
		ret.TreeWeight[i] = 1
	}
	return ret
}

func (uf *LUnionFind) GetSize() int {
	return uf.Size
}

func (uf *LUnionFind) find(p int) int {
	//find the id of elemnt p
	if p == uf.ParentId[p] {
		return p
	}
	for p != uf.ParentId[p] {
		p = uf.ParentId[p]
	}
	return p
}
func (uf *LUnionFind) IsConnnected(p, q int) bool {
	pid := uf.find(p)
	qid := uf.find(q)
	return pid == qid
}
func (uf *LUnionFind) max(a, b int) int {
	ret := a
	if b > a {
		ret = b
	}
	return ret
}
func (uf *LUnionFind) Union(p, q int) {
	//merge set of p and set of q
	pid := uf.find(p)
	qid := uf.find(q)
	if pid == qid {
		return
	}
	//here just change uf.ParentId[pid] = qid or uf.ParentId[qid] = pid according to the depth of rooted-tree-pid and rooted-tree-qid
	if uf.TreeWeight[pid] <= uf.TreeWeight[qid] {
		uf.ParentId[pid] = qid
		uf.TreeWeight[qid] = uf.max(uf.TreeWeight[qid], 1+uf.TreeWeight[pid])
	} else {
		uf.ParentId[qid] = pid
		uf.TreeWeight[pid] = uf.max(uf.TreeWeight[pid], 1+uf.TreeWeight[qid])
	}
	return
}
