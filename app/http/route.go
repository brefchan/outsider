package http

import (
	"github.com/bref/outsider/app/http/module/demo"
	"github.com/bref/outsider/framework/gin"
)

func Routes(r *gin.Engine) {
	r.Static("/dist/", "./dist/")

	demo.Register(r)
}
