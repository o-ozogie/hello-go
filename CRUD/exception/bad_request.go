package exception

import "github.com/valyala/fasthttp"

func InvalidParameterException() Exception {
	return Exception{errorCode: invalidParameter, message: "invalid parameter", status: fasthttp.StatusBadRequest}
}


