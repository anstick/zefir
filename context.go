package zefir

import "net/http"

type Context struct {
	req 	*Request
	res 	*Response
	route 	*Route
}

func NewContext() Context {
	return Context{}
}

func (c *Context) SetRequest(req *http.Request) {
	request := NewRequest(req)
	c.req  = &request
}

func (c *Context) GetRequest() *Request {
	return  c.req
}

func (c *Context) SetRoute(route *Route)  {
	c.route = route
}

func (c *Context) GetRoute() *Route {
	return  c.route
}

func (c *Context) SetResponse(res http.ResponseWriter) {
	response := NewResponse(res)
	c.res = &response
}

func (c *Context) GetResponse() *Response {
	return c.res
}

