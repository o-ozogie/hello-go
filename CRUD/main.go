package main

import (
	"fmt"
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
	"hello-go/CRUD/entity"
	routing "hello-go/CRUD/router"
)

func main() {
	r := router.New()

	entity.New().Connect()

	routing.UserRouter(r)

	fmt.Println("running at 8080")
	_ = fasthttp.ListenAndServe(":8080", r.Handler)
}
