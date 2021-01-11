package request

import (
	"encoding/json"
	"github.com/valyala/fasthttp"
	"hello-go/CRUD/exception"
)

type Request interface {
	Unmarshal(ctx *fasthttp.RequestCtx)
}

func unmarshal(ctx *fasthttp.RequestCtx, request Request) {
	if json.Unmarshal(ctx.Request.Body(), request) != nil {
		exception.RaiseException(ctx, exception.InvalidParameterException())
	}
}
