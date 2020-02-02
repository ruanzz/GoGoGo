package main

import "fmt"

/**
 * 函数作为返回值，Java中这种写法叫闭包
 */
func adder() func(int) int {
	result := 0
	return func(i int) int {
		result += i
		return result
	}
}

/**
 * 斐波那契数列
 */
func fib() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

func main() {
	// 函数作为变量
	result := adder()
	for i := 1; i < 10; i++ {
		fmt.Println(result(i))
	}

	fmt.Println("=========")
	f := fib()
	for i := 1; i < 10; i++ {
		fmt.Println(f())
	}

}
