package middleware

import (
	"encoding/json"
	"github.com/valyala/fasthttp"
	"hello-go/CRUD/request"
)

func UnMarshal(ctx *fasthttp.RequestCtx, t request.Request) bool {
	e := json.Unmarshal(ctx.Request.Body(), t)
	return e == nil || t.Validate()
}
