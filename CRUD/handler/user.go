package handler

import (
	"github.com/valyala/fasthttp"
	"hello-go/CRUD/entity"
	"hello-go/CRUD/exception"
	"hello-go/CRUD/middleware"
	"hello-go/CRUD/request"
)

func SignUpHandler(ctx *fasthttp.RequestCtx) {
	signUpRequest := new(request.SignUpRequest)

	if !middleware.UnMarshal(ctx, signUpRequest) {
		exception.RaiseException(ctx, exception.InvalidParameterException)
		return
	}

	if !SignUp(*signUpRequest) {
		exception.RaiseException(ctx, exception.ExistingEmailException)
		return
	}
}

func SignUp(signUpRequest request.SignUpRequest) bool {
	return entity.CreateUser(signUpRequest).RowsAffected == 1
}
