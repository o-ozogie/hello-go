package middleware

import (
	"github.com/valyala/fasthttp"
)

type Middleware func(ctx *fasthttp.RequestCtx) error

func FilterGroup(middlewares []Middleware) func(handler fasthttp.RequestHandler) fasthttp.RequestHandler {
	return func(handler fasthttp.RequestHandler) fasthttp.RequestHandler {
		return Filter(middlewares, handler)
	}
}

func Filter(middlewares []Middleware, handler fasthttp.RequestHandler) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		for _, middleware := range middlewares {
			e := middleware(ctx)
			if e != nil {
				return
			}
		}
		handler(ctx)
	}
}
