package main

import (
	"auth-app/db"
	"auth-app/handlers"
	"auth-app/middleware"
	"auth-app/ws"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	db.Init()

	r := gin.Default()
	r.Use(cors.Default()) // Allow frontend

	hub := ws.NewHub()
	go hub.Run()

	r.LoadHTMLGlob("static/*")

	r.GET("/signup", func(c *gin.Context) {
		c.HTML(200, "signup.html", nil)
	})
	r.GET("/login", func(c *gin.Context) {
		c.HTML(200, "login.html", nil)
	})

	r.POST("/signup", handlers.Signup)
	r.POST("/login", handlers.Login)

	// ✅ PUBLIC WebSocket route — keep this outside /api group!
	r.GET("/ws",
		ws.ServeWs(hub),
	)

	// Protected routes
	auth := r.Group("/api")
	auth.Use(middleware.JWTAuthMiddleware())
	{
		auth.GET("/me", handlers.GetMe)
		auth.GET("/users", handlers.ListUsers)
	}

	r.Run(":8080")
}
