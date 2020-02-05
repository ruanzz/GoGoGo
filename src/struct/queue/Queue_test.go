package queue

import "fmt"

// Queue.Pop 测试用例
func ExampleQueue_Pop() {
	q := Queue{1}
	q.Push(2)
	fmt.Println(q.Pop())

	// Output:
	// 1
}
