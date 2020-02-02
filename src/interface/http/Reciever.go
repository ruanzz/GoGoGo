package http

import (
	"net/http"
	"net/http/httputil"
)

type Reciever struct {
	UserAgent string
}

func (r Reciever) Get(url string) string {

	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	result, err := httputil.DumpResponse(resp, true)
	if err != nil {
		panic(err)
	}
	return string(result)
}
