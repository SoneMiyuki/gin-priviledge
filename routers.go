package main

import (
	"github.com/gin-gonic/gin"
	"priviledge/Middleware"
	"priviledge/controller"
)

func CollectRouter(r *gin.Engine) *gin.Engine {
	r.POST("/api/auth/register", controller.Register)
	r.POST("/api/auth/login", controller.Login)
	r.GET("/api/auth/info", Middleware.AuthMiddleware(), controller.Info)
	return r
}
