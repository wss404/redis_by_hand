package hashtable

type HNode struct {
	Next  *HNode
	HCode uint64
}

type HTab struct {
	tab  []*HNode
	mask uint64 // hashcode长度最大为64位
	size uint64
}

func InitHashTable(n uint64) *HTab {
	if n < 0 || (n-1)&n != 0 {
		panic("illegal table size.")
	}
	h := &HTab{}
	h.tab = make([]*HNode, n)
	h.mask = n - 1
	return h
}

func (h *HTab) Insert(node *HNode) {
	pos := node.HCode & h.mask
	next := h.tab[pos]
	node.Next = next
	h.tab[pos] = node
	h.size++
}

func (h *HTab) LookUp(key *HNode, cmp func(*HNode, *HNode) bool) **HNode {
	if h.tab == nil {
		return nil
	}

	pos := key.HCode & h.mask
	// 找到【指向目标节点的指针的地址】，它可以是目标节点上一节点的地址，也可以是链表的头节点，便于后续的删除
	from := &h.tab[pos]

	for *from != nil {
		if cmp(*from, key) {
			return from
		}
		from = &(*from).Next
	}

	return nil
}

// Detach 从单链表中删除一个节点
func (h *HTab) Detach(from **HNode) *HNode {
	node := *from
	*from = (*from).Next
	h.size--
	return node
}

func (h *HTab) Scan(f func(*HNode, *[]byte), out *[]byte) {
	if h.size == 0 {
		return
	}
	for i := uint64(0); i <= h.mask; i++ {
		node := h.tab[i]
		for node != nil {
			f(node, out)
			node = node.Next
		}
	}
}

func HNodeSame(lhs, rhs *HNode) bool {
	return lhs == rhs
}
