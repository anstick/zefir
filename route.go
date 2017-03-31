package zefir

import "regexp"

type Handler func(context Context)

type Route struct {
	method string
	name string
	pattern *regexp.Regexp
	handler Handler
}

func (r *Route) GetHandler() Handler  {
	return r.handler
}
