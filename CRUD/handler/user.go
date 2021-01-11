package handler

import (
	"github.com/valyala/fasthttp"
	"hello-go/CRUD/entity"
	"hello-go/CRUD/exception"
	"hello-go/CRUD/request"
)

func SignUpHandler(ctx *fasthttp.RequestCtx) {
	signUpRequest := request.SignUpRequest{}
	signUpRequest.Unmarshal(ctx)
	signUp(signUpRequest)
}

func signUp(signUpRequest request.SignUpRequest) {
	if entity.CreateUser(signUpRequest).RowsAffected != 1 {
		exception.Raise(exception.ExistingEmailException())
	}
}
