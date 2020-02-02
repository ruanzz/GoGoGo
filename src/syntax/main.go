package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"math"
	"math/cmplx"
	"os"
	"strconv"
)

var (
	k1 = 1
	k2 = "11"
)

func sayHi(a string) {
	var hello = "Hello "
	fmt.Println(hello, a)
}

func varDeduction() {
	var numA, stringB = 1, "11"
	fmt.Println(numA, stringB)

}

func complex() {
	c := 3 + 4i
	fmt.Println(cmplx.Abs(c))
}

func euler() {
	fmt.Printf("%.3f \n", cmplx.Pow(math.E, 1i*math.Pi)+1)
}

func explicitCast() {
	var a, b = 3, 4
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}

func enum() {
	const (
		java = iota * 2 // iota说明是从0开始，后边开始递增1,并且后边是代入表达式进行计算
		python
		_ // 占位符
		golang
		javascript
	)
	fmt.Println(java, python, golang, javascript)

}

func branch() {
	const file = "src/syntax/text.out"
	if bytes, e := ioutil.ReadFile(file); e != nil {
		fmt.Println(e)
	} else {
		fmt.Println(bytes)
	}
}

func convert2Bin(n int) string {
	result := ""
	for ; n > 0; n = n / 2 {
		result += strconv.Itoa(n % 2)
	}
	fmt.Println(result)
	return result
}

func readFile(filename string) string {
	result := ""
	if file, e := os.Open(filename); e == nil {
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			result += fmt.Sprintf("%s \n", scanner.Text())
		}
	}
	fmt.Println(result)
	return result
}

func sum(n ...int) int {
	result := 0
	for i := range n {
		result += i
	}
	fmt.Println(result)
	return result
}

func swap(a, b *int) (int, int) {
	fmt.Println(*a, *b)
	*a, *b = *b, *a
	fmt.Println(*a, *b)
	return *a, *b
}

func main() {
	lang := "Go "
	sayHi(lang)
	varDeduction()
	fmt.Println(k1, k2)
	complex()
	euler()
	explicitCast()
	enum()
	branch()
	convert2Bin(11)
	const filename = "src/syntax/text.out"
	readFile(filename)
	sum(1, 2, 3, 4, 5)

	a, b := 1, 2
	c, d := swap(&a, &b)
	fmt.Println(a, b, c, d)

	// array
	var arr [5]int
	fmt.Println("arr:", arr)
	arr1 := [6]int{1, 2, 3, 4, 5, 6}
	fmt.Println("arr1:", arr1)

	for i, v := range arr1 {
		fmt.Println(i, v)
	}

	// slice
	slice := arr1[3:5]
	fmt.Println(slice)
	slice[0] = 100
	fmt.Println("slice:", slice)
	fmt.Println("arr1:", arr1)

	slice1 := append(slice, 101)
	fmt.Println(slice1)
	fmt.Println(slice)
	fmt.Println(arr1)

	// map
	m := map[string]string{
		"key1": "value1",
		"key2": "value2",
	}
	fmt.Println(m)
	for k, v := range m {
		fmt.Println(k, v)
	}

	m1 := make(map[string]int)
	fmt.Println(m1)

	if value1, ok := m["key1"]; ok {
		fmt.Println(value1)
	} else {
		fmt.Println("map not contain key")
	}

	delete(m, "key1")
	if value1, ok := m["key1"]; ok {
		fmt.Println(value1)
	} else {
		fmt.Println("map not contain key")
	}

}
