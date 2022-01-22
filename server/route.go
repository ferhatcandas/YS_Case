package server

import (
	"strings"
)

type Route struct {
	Pattern string
	Subset  []Subset
	Add     func(realPath, method string, f func(res Response, req Request) *Response)
	Search  func(realPath, method string) (*Subset, func(res Response, req Request) *Response)
}
type Subset struct {
	RealPath string
	Params   map[int]string
	Function map[string]func(res Response, req Request) *Response
}

func NewRoute(path, method string, f func(res Response, req Request) *Response) Route {
	pattern := strings.Split(path, ":")[0]
	pattern = strings.TrimSuffix(pattern, "/")
	r := Route{Pattern: pattern}
	r.Add = r.Push
	r.Search = r.Find
	r.Add(path, method, f)
	return r
}

func (r *Route) Find(realPath, method string) (*Subset, func(res Response, req Request) *Response) {
	params := getParams(r.Pattern, realPath)
	for _, value := range r.Subset {
		if value.Function[method] != nil {
			if len(value.Params) == len(params) {
				for i, item := range params {
					value.Params[i] = item
				}
				return &value, value.Function[method]
			}
		}
	}
	return nil, nil
}
func getParams(pattern, realPath string) map[int]string {
	pattern = strings.TrimPrefix(pattern, "/")
	paths := strings.Split(realPath, "/")
	params := make(map[int]string)
	index := 0
	for _, value := range paths {
		if value != pattern && value != "" {
			params[index] = value
			index++
		}
	}
	return params
}

func (r *Route) Push(realPath, method string, f func(res Response, req Request) *Response) {

	st := struct {
		RealPath string
		Params   map[int]string
		Function map[string]func(res Response, req Request) *Response
	}{
		RealPath: realPath,
		Params:   make(map[int]string),
		Function: make(map[string]func(res Response, req Request) *Response),
	}
	st.Function[method] = f
	st.Params = getParams(r.Pattern, realPath)
	r.Subset = append(r.Subset, st)
}
