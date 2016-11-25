package main

import (
	"log"

	"github.com/valyala/fasthttp"
)

var proxyClient = &fasthttp.HostClient{
	Addr: "10.0.1.2:8080",
}

func ReverseProxyHandler(ctx *fasthttp.RequestCtx) {
	req := &ctx.Request
	resp := &ctx.Response
	req.Header.Del("Connection")
	if err := proxyClient.Do(req, resp); err != nil {
		ctx.Logger().Printf("error when proxying the request: %s", err)
	}
	resp.Header.Del("Connection")
}
func main() {
	if err := fasthttp.ListenAndServe(":80", reverseProxyHandler); err != nil {
		log.Fatalf("error in fasthttp server: %s", err)
	}
}
