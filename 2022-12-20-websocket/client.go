package main

import (
	"fmt"
	"log"

	"github.com/gorilla/websocket"
)

func main() {
	conn, _, err := websocket.DefaultDialer.Dial("ws://127.0.0.1:8088/ws", nil)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	for {
		fmt.Println("Writing the message")
		// Send message to server
		err := conn.WriteMessage(websocket.TextMessage, []byte("Hello from client!"))
		if err != nil {
			log.Println(err)
			break
		}

		fmt.Println("Waiting for the server message")
		// Read message from server
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			break
		}
		fmt.Printf("Received message: %s\n", message)
	}
}

