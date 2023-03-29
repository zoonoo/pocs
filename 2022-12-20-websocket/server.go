package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

// Connection represents a WebSocket connection
type Connection struct {
	ID   string
	Conn *websocket.Conn
}

// Server is a WebSocket server
type Server struct {
	Connections map[string]*Connection
}

// NewServer creates a new Server instance
func NewServer() *Server {
	return &Server{
		Connections: make(map[string]*Connection),
	}
}

// AddConnection adds a new WebSocket connection to the server
func (s *Server) AddConnection(conn *websocket.Conn) {
	c := &Connection{
		ID:   fmt.Sprintf("conn-%d", len(s.Connections)+1),
		Conn: conn,
	}
	s.Connections[c.ID] = c
}

// RemoveConnection removes a WebSocket connection from the server
func (s *Server) RemoveConnection(conn *websocket.Conn) {
	for id, c := range s.Connections {
		if c.Conn == conn {
			delete(s.Connections, id)
			break
		}
	}
}

func main() {
	server := NewServer()

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			return
		}
		defer conn.Close()

		server.AddConnection(conn)
		defer server.RemoveConnection(conn)

		for {
			// Read message from client
			_, message, err := conn.ReadMessage()
			if err != nil {
				break
			}
			fmt.Println("message received.")

			// Send message to all other connections
			for _, c := range server.Connections {
				if c.Conn != conn {
					c.Conn.WriteMessage(websocket.TextMessage, message)
				}
			}
		}
	})

	http.ListenAndServe(":8088", nil)
}

