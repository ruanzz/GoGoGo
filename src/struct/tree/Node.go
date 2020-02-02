package tree

import (
	"fmt"
)

/**
 * 定义一个结构体以及变量
 */
type Node struct {
	Value int
}

/**
 * 定义结构体的方法
 * 注意这里的node是传值的不是传引用，Go语言中要想传引用就传指针
 */
func (node Node) Printf() {
	fmt.Println(node.Value)
}

/**
 * 传引用设值
 */
func (node *Node) SetValue(value int) {
	node.Value = value
}
