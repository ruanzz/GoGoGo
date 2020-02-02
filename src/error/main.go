package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

/**
 * 写文件
 */
func writeFile(fileName string) {
	file, e := os.Create(fileName)
	if e != nil {
		panic(e)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()
	for i := 1; i <= 100; i++ {
		fmt.Fprintf(writer, "i: %d\n", i)
	}
}

/**
 * defer匿名函数捕捉错误
 */
func tryRecover() {

	defer func() {
		r := recover()
		if err, ok := r.(error); ok {
			fmt.Println("Error occurred: ", err)
		} else {
			panic(err)
		}
	}()

	// 抛出一个错误
	panic(errors.New("Invalid"))
}

func main() {
	writeFile("src/error/log.out")

	tryRecover()

}
