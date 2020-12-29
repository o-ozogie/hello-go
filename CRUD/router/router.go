package router

import (
	"github.com/fasthttp/router"
	"hello-go/CRUD/handler"
)

func UserRouter(r *router.Router) {
	userRouter := r.Group("/user")

	userRouter.POST("/signup", handler.SignUpHandler)
}
