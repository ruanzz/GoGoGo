package handler

import (
	"error/web/common"
	"net/http"
	"strings"
)

/**
 * Api 路由
 */
func HandlerApi(writer http.ResponseWriter, request *http.Request) error {
	path := request.URL.Path
	if strings.HasPrefix(path, FileApi) {
		return HandleFileApi(writer, request)
	}
	return common.AppError{Message: "No route"}
}
