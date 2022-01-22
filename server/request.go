package server

import (
	"io/ioutil"
	"net/http"
)

type Request struct {
	r        *http.Request
	Params   map[int]string
	Header   map[string][]string
	Body     []byte
	RemoteIP string
	Method   string
	URL      string
}

func NewRequest(path string, params map[int]string, r *http.Request) Request {
	// params := getParams(path, r)
	body, _ := ioutil.ReadAll(r.Body)
	headers := make(map[string][]string)
	for name, values := range r.Header {
		for _, value := range values {
			headers[name] = append(headers[name], value)
			break
		}
	}
	return Request{r: r, Params: params, Body: body, Header: r.Header, RemoteIP: r.RemoteAddr, Method: r.Method, URL: r.Host + r.URL.Path}
}
