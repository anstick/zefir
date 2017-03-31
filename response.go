package zefir

import (
	"net/http"
)

type Response struct {
	http.ResponseWriter
	statusCode int
}

func NewResponse(resp http.ResponseWriter) Response {
	return Response{resp, 200}
}


func (r *Response) SetStatusCode(statusCode int) {
	r.WriteHeader(statusCode)
	r.statusCode = statusCode
}

func (r *Response) GetStatusCode() int {
	return r.statusCode
}

func (r *Response) SetText(t string) {
	//TODO Shitty realisation
	r.Write([]byte(t))
}
