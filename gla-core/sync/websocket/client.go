package websocket

import (
	"fmt"
	"log"
	"time"

	"github.com/gorilla/websocket"
)

type WebSocketClient struct {
	conn *websocket.Conn
	url  string
}

func NewClient(url string) (*WebSocketClient, error) {
	conn, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		return nil, fmt.Errorf("error dialing WebSocket server: %v", err)
	}
	return &WebSocketClient{
		conn: conn,
		url:  url,
	}, nil
}

func (ws *WebSocketClient) SendMessage(message string) error {
	err := ws.conn.WriteMessage(websocket.TextMessage, []byte(message))
	if err != nil {
		return fmt.Errorf("error sending message: %v", err)
	}
	return nil
}

func (ws *WebSocketClient) ReceiveMessage() (string, error) {
	_, message, err := ws.conn.ReadMessage()
	if err != nil {
		return "", fmt.Errorf("error receiving message: %v", err)
	}
	return string(message), nil
}

func (ws *WebSocketClient) Close() error {
	err := ws.conn.Close()
	if err != nil {
		return fmt.Errorf("error closing WebSocket connection: %v", err)
	}
	return nil
}
