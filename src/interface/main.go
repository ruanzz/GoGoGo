package main

import (
	"fmt"
	"interface/http"
	"interface/mock"
)

type Reciever interface {
	Get(url string) string
}

/**
 * 下载
 */
func download(r Reciever) string {
	return r.Get("http://192.168.56.5:30315")
}

func main() {
	fmt.Println("========")
	var r Reciever = mock.Reciever{}
	fmt.Println(r)
	fmt.Println(download(r))

	fmt.Println("=========")
	r = http.Reciever{}
	fmt.Println(download(r))

	fmt.Println("========")
	switch r.(type) {
	case http.Reciever:
		fmt.Println("http.Reciever")
	case mock.Reciever:
		fmt.Println("mock.Reciever")

	}
	if reciever, ok := r.(mock.Reciever); ok {
		fmt.Println("mock.Reciever Content:", reciever.Content)
	}

	if reciever, ok := r.(http.Reciever); ok {
		fmt.Println("http.Reciever UserAgent:", reciever.UserAgent)
	}

}
