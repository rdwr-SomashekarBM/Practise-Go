package main

import (
	"auth-app/db"
	"auth-app/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	db.Init()

	r := gin.Default()
	r.LoadHTMLGlob("static/*")

	r.GET("/signup", func(c *gin.Context) {
		c.HTML(200, "signup.html", nil)
	})
	r.GET("/login", func(c *gin.Context) {
		c.HTML(200, "login.html", nil)
	})

	r.POST("/signup", handlers.Signup)
	r.POST("/login", handlers.Login)

	r.Run(":8080")
}
