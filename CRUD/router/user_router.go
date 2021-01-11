package router

import (
	"github.com/fasthttp/router"
	"hello-go/CRUD/handler"
	"hello-go/CRUD/middleware"
)

func UserRouter(r *router.Router) {
	userRouter := r.Group("/user")

	userRouter.POST("/signup", middleware.Filter([]middleware.Middleware{}, handler.SignUpHandler))
}
