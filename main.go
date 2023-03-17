package main

import (
	"github.com/gin-gonic/gin"
	"github.com/khatrisaugat/JWTPractise/controllers"
	"github.com/khatrisaugat/JWTPractise/middlewares"
)

func main() {
	router := gin.Default()
	public := router.Group("/api")
	public.POST("/register", controllers.Register)
	public.POST("/login", controllers.Login)
	protected := router.Group("/api/admin")
	protected.Use(middlewares.JwtAuthMiddleware())
	protected.GET("/user", controllers.CurrentUser)
	router.Run(":5000")
}
