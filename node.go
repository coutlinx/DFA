package daf

// DFANode 节点结构体
type DFANode struct {
	Children map[rune]*DFANode
	isEnd    bool
}

func newNode() *DFANode {
	return &DFANode{Children: make(map[rune]*DFANode)}
}
