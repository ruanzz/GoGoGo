package handler

import (
	"error/web/common"
	"io/ioutil"
	"net/http"
	"os"
)

const (
	FileApi = "/files/"
)

/**
 * 处理文件Api相关业务逻辑
 */
func HandleFileApi(writer http.ResponseWriter, request *http.Request) error {
	path := request.URL.Path
	fileName := path[len(FileApi):]
	root := "src/error/"
	var filePath = root + fileName
	file, e := os.Open(filePath)
	if e != nil {
		return common.AppError{Message: "Can't open file " + fileName, ErrorInfo: e}
	}
	defer file.Close()

	bytes, err := ioutil.ReadAll(file)

	if err != nil {
		return common.AppError{Message: "Read file error", ErrorInfo: err.Error()}
	}
	writer.Write(bytes)
	return nil
}
