Zefir - A Go micro web-framework
================================
Inspired by Django Project

Installation
------------
```bash
$ cd $GOPATH/src
$ git clone git@github.com:anstick/zefir.git github.com/anstick/zefir
```


A simple application example
----------------------------
```go
package main

import (
	"github.com/anstick/zefir"
)

func main() {
	app := zefir.Create()

	router:= app.GetRouter()
	router.Get("index", "/$", func(context zefir.Context) {
		context.GetResponse().SetText("Hello world")
	})

	app.Start()
}
```