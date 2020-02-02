package mock

import "fmt"

/**
 * mock 接口实现
 */
type Reciever struct {
	Content string
}

/**
 * 类似Java中toString(),Stringer接口中的方法
 */
func (r Reciever) String() string {
	return fmt.Sprintf("Reciever: {Content=%s}", r.Content)
}

func (r Reciever) Get(url string) string {
	return r.Content
}
