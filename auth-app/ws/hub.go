package ws

import (
	"github.com/gorilla/websocket"
)

type Client struct {
	ID   int
	Conn *websocket.Conn
	Send chan Message
	Hub  *Hub
}

type Hub struct {
	Clients    map[int]*Client // map of userID -> Client
	Broadcast  chan Message
	Register   chan *Client
	Unregister chan *Client
}

type Message struct {
	SenderID   int    `json:"sender_id"`
	ReceiverID int    `json:"receiver_id"`
	Content    string `json:"content"`
}

func NewHub() *Hub {
	return &Hub{
		Clients:    make(map[int]*Client),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.Register:
			h.Clients[client.ID] = client

		case client := <-h.Unregister:
			if _, ok := h.Clients[client.ID]; ok {
				delete(h.Clients, client.ID)
				close(client.Send)
			}

		case msg := <-h.Broadcast:
			// Send only to receiver
			if receiver, ok := h.Clients[msg.ReceiverID]; ok {
				receiver.Send <- msg
			}
		}
	}
}
