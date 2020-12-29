package exception

import "github.com/valyala/fasthttp"

func InvalidParameterException() *exception {
	return &exception{errorCode: invalidParameter, message: "invalid parameter", status: fasthttp.StatusBadRequest}
}


