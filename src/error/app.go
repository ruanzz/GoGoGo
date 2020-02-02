package main

import (
	"error/web/common"
	"error/web/handler"
	"fmt"
	"log"
	"net/http"
)

// 定义handler
type appHandler func(writer http.ResponseWriter, request *http.Request) error

/**
 * 统一错误处理
 */
func errWrapper(handler appHandler) func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		e := handler(writer, request)
		if e == nil {
			return
		}
		message := http.StatusText(http.StatusInternalServerError)
		if err, ok := e.(common.AppError); ok {
			message = err.Message
		}
		log.Println(e)
		http.Error(writer, message, http.StatusInternalServerError)

	}
}

func main() {

	http.HandleFunc("/", errWrapper(handler.HandlerApi))

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Web applicaton start failed :", err)
		return
	}

	fmt.Println("Web applicaton started")
}
