package exception

import "github.com/valyala/fasthttp"

func ExistingEmailException() *exception {
	return &exception{errorCode: existingEmail, message: "already existing email", status: fasthttp.StatusConflict}
}
