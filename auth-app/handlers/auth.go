package handlers

import (
	"auth-app/db"
	"auth-app/models"
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Signup(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	_, err := db.DB.Exec("INSERT INTO users (username, password) VALUES ($1, $2)", username, string(hash))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User exists or invalid input"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Signup successful"})
}

func Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	var user models.User
	err := db.DB.QueryRow("SELECT id, password FROM users WHERE username=$1", username).Scan(&user.ID, &user.Password)
	if err == sql.ErrNoRows {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
}
