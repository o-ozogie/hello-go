package exception

import (
	"fmt"
	"github.com/valyala/fasthttp"
	"log"
)

const(
	invalidParameter = "X0400"
	existingEmail = "X0409"
)

type exception struct {
	errorCode string
	message   string
	status    int
}

func (e exception) Error() string {
	return fmt.Sprintf("[%s] %s", e.errorCode, e.message)
}

func Raise(e error) {
	log.Panic(e)
}

func RaiseException(ctx *fasthttp.RequestCtx, f func() *exception) {
	e := f()
	ctx.Error(e.Error(), e.status)
}
