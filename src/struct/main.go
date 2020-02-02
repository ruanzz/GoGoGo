package main

import (
	"fmt"
	"struct/queue"
	"struct/tree"
)

func main() {
	fmt.Println("===========")
	// 初始化实例
	node := &tree.Node{Value: 1}
	node.SetValue(7)
	node.Printf()

	fmt.Println("===========")
	// 二叉树
	treeNode := tree.BinaryTree{
		Root: *node,
	}
	treeNode.Right = &tree.BinaryTree{Root: tree.Node{6}}
	treeNode.Left = &tree.BinaryTree{Root: tree.Node{3}}

	treeNode.Left.Right = &tree.BinaryTree{Root: tree.Node{2}}
	treeNode.Left.Left = &tree.BinaryTree{Root: tree.Node{1}}

	treeNode.Right.Right = &tree.BinaryTree{Root: tree.Node{5}}
	treeNode.Right.Left = &tree.BinaryTree{Root: tree.Node{4}}
	treeNode.PostOrder()

	fmt.Println("=========")
	// 包装队列
	queue := queue.Queue{}
	queue.Push(1)
	queue.Push(2)
	fmt.Println(queue.Pop())
	fmt.Println(queue.Pop())
	fmt.Println(queue.IsEmpty())
	//fmt.Println(queue.Pop())
}
