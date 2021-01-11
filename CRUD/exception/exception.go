package exception

import (
	"fmt"
	"github.com/valyala/fasthttp"
	"log"
	"strconv"
)

const(
	invalidParameter = "X0400"
	existingEmail = "X0409"
	internalServerError = "E0500"
)

type Exception struct {
	errorCode string
	message   string
	status    int
}

func (e Exception) Error() string {
	return fmt.Sprintf("[%s] %s", e.errorCode, e.message)
}

func ToException(errorString string) (exception Exception, err error) {
	exception.errorCode = errorString[1:6]
	exception.message = errorString[8:]
	exception.status, err = strconv.Atoi(exception.errorCode[2:])

	return
}

func Raise(e Exception) {
	log.Panic(e)
}

func RaiseException(ctx *fasthttp.RequestCtx, e Exception) {
	ctx.Error(e.Error(), e.status)
}
