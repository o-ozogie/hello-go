package exception

import "github.com/valyala/fasthttp"

func ExistingEmailException() Exception {
	return Exception{errorCode: existingEmail, message: "already existing email", status: fasthttp.StatusConflict}
}
