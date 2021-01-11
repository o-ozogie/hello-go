package response

import (
	"encoding/json"
	"fmt"
	"github.com/valyala/fasthttp"
	"log"
)

type Response interface {
	Marshal(*fasthttp.RequestCtx, int)
}

func marshal(ctx *fasthttp.RequestCtx, statusCode int, response Response) {
	fmt.Println(response)
	marshaledResponse, e := json.Marshal(response)
	if e != nil {
		log.Panic(e)
	}
	ctx.Response.SetStatusCode(statusCode)
	ctx.Response.SetBody(marshaledResponse)
}
