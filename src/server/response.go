package server

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	writer http.ResponseWriter
	Status int
	Body   []byte
}

func NewResponse(wr http.ResponseWriter) Response {
	return Response{writer: wr}
}

func (res *Response) JSON(x interface{}, statusCode int) *Response {
	res.writer.WriteHeader(statusCode)
	res.Status = statusCode
	r, err := json.Marshal(x)
	if err != nil {
		res.writer.WriteHeader(http.StatusInternalServerError)
		res.writer.Write([]byte(err.Error()))
	}
	res.Body = r
	res.writer.Header().Add("Content-Type", "application/json")
	res.writer.Write(r)
	return res
}
