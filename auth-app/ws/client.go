package ws

import (
	"auth-app/db"
	"log"
)

// ReadPump handles incoming messages from the WebSocket connection
func (c *Client) ReadPump() {
	defer func() {
		c.Hub.Unregister <- c // Unregister client from hub
		c.Conn.Close()        // Close WebSocket connection
	}()

	for {
		var msg Message

		// Read message from WebSocket
		if err := c.Conn.ReadJSON(&msg); err != nil {
			log.Println("read error:", err)
			break
		}

		// Attach sender ID to the message
		msg.SenderID = c.ID

		// Save message to DB
		_, err := db.DB.Exec(
			"INSERT INTO messages (sender_id, receiver_id, content) VALUES ($1, $2, $3)",
			msg.SenderID,
			msg.ReceiverID,
			msg.Content,
		)
		if err != nil {
			log.Println("db insert error:", err)
			continue
		}

		// Send message to the intended receiver if they're connected
		if receiverClient, ok := c.Hub.Clients[msg.ReceiverID]; ok {
			receiverClient.Send <- msg
		}
	}
}

// WritePump handles outgoing messages to the WebSocket connection
func (c *Client) WritePump() {
	defer c.Conn.Close()

	for msg := range c.Send {
		if err := c.Conn.WriteJSON(msg); err != nil {
			log.Println("write error:", err)
			break
		}
	}
}
