package ws

import (
	"auth-app/utils"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // Allow all origins (adjust as needed)
	},
}

// ✅ THIS is the route handler used in main.go
func ServeWs(hub *Hub) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Query("token")
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing token"})
			return
		}

		userID, err := utils.ParseJWT(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}

		// 🔁 Reuse ServeWsHandler below
		ServeWsHandler(hub, c.Writer, c.Request, userID)
	}
}

// ✅ This is the actual WebSocket logic reused by above
func ServeWsHandler(hub *Hub, w http.ResponseWriter, r *http.Request, userID int) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("WebSocket upgrade error:", err)
		return
	}

	client := &Client{
		Hub:  hub,
		Conn: conn,
		Send: make(chan Message, 256),
		ID:   userID, // ✅ Comes from token
	}

	hub.Register <- client

	go client.WritePump()
	go client.ReadPump()
}
