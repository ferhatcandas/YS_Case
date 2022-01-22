package server

import (
	"fmt"
	"net/http"
)

type Server struct {
	mux        *http.ServeMux
	middleware func(res Response, req Request) *Response
	router     Router
}

func NewServer() Server {
	mux := http.NewServeMux()
	return Server{mux: mux, router: NewRouter()}
}
func (s *Server) Middleware(f func(res Response, req Request) *Response) {
	s.middleware = f
}
func (s *Server) handleRoute(path string, method string, f func(res Response, req Request) *Response) {
	s.router.Add(path, method, f)
}
func (s *Server) GET(path string, f func(res Response, req Request) *Response) {
	s.handleRoute(path, http.MethodGet, f)
}

func (s *Server) POST(path string, f func(res Response, req Request) *Response) {
	s.handleRoute(path, http.MethodPost, f)
}
func (s *Server) DELETE(path string, f func(res Response, req Request) *Response) {
	s.handleRoute(path, http.MethodDelete, f)
}

func (s *Server) Run(port string) {
	fmt.Println("Server started at port http://localhost" + port)

	s.mux.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {

		subset, fun := s.router.Find(r.URL.Path, r.Method)
		if subset != nil {
			res := NewResponse(rw)
			req := NewRequest(subset.RealPath, subset.Params, r)
			fun(res, req)
			if s.middleware != nil {
				s.middleware(res, req)
			}

		} else {
			rw.WriteHeader(http.StatusNotFound)
		}

	})

	err := http.ListenAndServe(port, s.mux)
	if err != nil {
		panic(err)
	}
}
