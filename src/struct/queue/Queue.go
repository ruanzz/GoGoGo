package queue

// 将slice定义为别名Queue
type Queue []interface{}

// 入队列
func (queue *Queue) Push(value interface{}) {
	*queue = append(*queue, value)
}

// 出队列
func (queue *Queue) Pop() interface{} {
	if queue.IsEmpty() {
		// 抛错
		panic("队列已空")
	}
	head := (*queue)[0]
	*queue = (*queue)[1:]
	return head
}

// 队列判空
func (queue *Queue) IsEmpty() bool {
	return len(*queue) == 0
}
