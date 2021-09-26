package main

import (
	"github.com/gin-gonic/gin"
	"priviledge/controller"
)

func CollectRouter(r *gin.Engine) *gin.Engine{
	r.POST("/api/auth/register", controller.Register)
	return r
}