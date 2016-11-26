package main

import (
	"log"

	"github.com/valyala/fasthttp"
)

// Hop-by-hop headers. These are removed when sent to the backend.
// http://www.w3.org/Protocols/rfc2616/rfc2616-sec13.html
var hopHeaders = []string{
	"Connection",
	"Keep-Alive",
	"Proxy-Authenticate",
	"Proxy-Authorization",
	"Te", // canonicalized version of "TE"
	"Trailers",
	"Transfer-Encoding",
	"Upgrade",
}

var proxyClient = &fasthttp.HostClient{
	Addr: "10.0.1.2:8080",
}

func removeReqHeader(header fasthttp.RequestHeader) {
	for _, h := range hopHeaders {
		header.Del(h)
	}
}

func removeResHeader(header fasthttp.ResponseHeader) {
	for _, h := range hopHeaders {
		header.Del(h)
	}
}

//
//func addForwardedFor(req *fasthttp.Request) {
//	if clientIP, _, err := net.SplitHostPort(req.RemoteAddr); err == nil {
//		if prior, ok := req.Header["X-Forwarded-For"]; ok {
//			clientIP = strings.Join(prior, ", ") + ", " + clientIP
//		}
//		req.Header.Set("X-Forwarded-For", clientIP)
//	}
//}

func reverseProxyHandler(ctx *fasthttp.RequestCtx) {
	req := &ctx.Request
	resp := &ctx.Response
	removeReqHeader(req.Header)
	//	addForwardedFor(req.Header)
	if err := proxyClient.Do(req, resp); err != nil {
		ctx.Logger().Printf("error when proxying the request: %s", err)
	}
	removeResHeader(resp.Header)
}

func main() {
	if err := fasthttp.ListenAndServe(":80", reverseProxyHandler); err != nil {
		log.Fatalf("error in fasthttp server: %s", err)
	}
}
