package request

import (
	"github.com/valyala/fasthttp"
)

type SignUpRequest struct {
	Email string `json:"email"`
	Password string `json:"password"`
	Name string `json:"name"`
}

func (request *SignUpRequest) Unmarshal(ctx *fasthttp.RequestCtx) {
	unmarshal(ctx, request)
}
