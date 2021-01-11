package exception

import "github.com/valyala/fasthttp"

func InternalServerError() Exception {
	return Exception{errorCode: internalServerError, message: "internal server error", status: fasthttp.StatusInternalServerError}
}
