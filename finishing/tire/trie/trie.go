package trie

type Node struct {
	children    map[rune]*Node
	Content     rune
	FullContent string
	End         bool
}

func (n *Node) AddChild(c rune) *Node {

	if n.children == nil {
		n.children = make(map[rune]*Node)
	}

	curNode, ok := n.children[c]
	if ok {
		return curNode
	}

	// 不存在c的敏感词节点

	n.children[c] = &Node{
		Content: c,
	}
	return n.children[c]

}

func (n *Node) FindChild(c rune) *Node {
	if len(n.children) == 0 {
		return nil
	}

	return n.children[c]
}
