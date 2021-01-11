package middleware

import (
	"github.com/valyala/fasthttp"
	"hello-go/CRUD/exception"
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
		defer ErrorHandle(ctx)
		handler(ctx)
	}
}

func ErrorHandle(ctx *fasthttp.RequestCtx) {
	if r := recover(); r != nil {
		x, e := exception.ToException(r.(string))
		if e != nil {
			exception.RaiseException(ctx, exception.InternalServerError())
			return
		}
		exception.RaiseException(ctx, x)
	}
}
