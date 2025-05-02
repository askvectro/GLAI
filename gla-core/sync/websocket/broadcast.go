package websocket

import (
	"fmt"
	"log"
	"time"
)

func (ws *WebSocketServer) BroadcastMessage(message string) {
	for client := range ws.clients {
		err := client.WriteMessage(1, []byte(message))
		if err != nil {
			log.Println("Error broadcasting message:", err)
			client.Close()
			delete(ws.clients, client)
		}
	}
}

func (ws *WebSocketServer) BroadcastPeriodically(message string, interval time.Duration) {
	for {
		ws.BroadcastMessage(message)
		time.Sleep(interval)
	}
}
