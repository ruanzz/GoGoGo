package tree

/**
 * 通过组合的方式达到继承目的
 */
type BinaryTree struct {
	Root        Node
	Left, Right *BinaryTree
}

/**
 * 后续遍历
 */
func (treeNode BinaryTree) PostOrder() {
	if treeNode.Left != nil {
		treeNode.Left.PostOrder()
	}
	if treeNode.Right != nil {
		treeNode.Right.PostOrder()
	}
	treeNode.Root.Printf()
}
