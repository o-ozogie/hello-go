package response

import "github.com/valyala/fasthttp"

type SuccessfulResponse struct {
	Message string `json:"message"`
}

func (response *SuccessfulResponse) Marshal(ctx *fasthttp.RequestCtx, statusCode int) {
	marshal(ctx, statusCode, response)
}
