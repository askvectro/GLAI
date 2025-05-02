package websocket

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/websocket"
)

type WebSocketServer struct {
	upgrader websocket.Upgrader
	clients  map[*websocket.Conn]bool
}

func NewServer() *WebSocketServer {
	return &WebSocketServer{
		upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool { return true },
		},
		clients: make(map[*websocket.Conn]bool),
	}
}

func (ws *WebSocketServer) HandleConnections(w http.ResponseWriter, r *http.Request) {
	conn, err := ws.upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Error upgrading connection:", err)
		return
	}
	defer conn.Close()

	ws.clients[conn] = true

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("Error reading message:", err)
			delete(ws.clients, conn)
			break
		}
		log.Printf("Received message: %s", message)
	}
}

func (ws *WebSocketServer) StartServer(addr string) {
	http.HandleFunc("/", ws.HandleConnections)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
