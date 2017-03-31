package zefir

import "net/http"

type Request struct {
	httpRequest *http.Request
}

func NewRequest(req *http.Request) Request {
	return Request{httpRequest:req}
}

func (r *Request) GetHttpRequest() *http.Request {
	return r.httpRequest
}