package sync

import (
	"log"
	"gla-core/sync/firebase"
	"gla-core/sync/websocket"
)

// SyncEngine holds the configuration for both Firebase and WebSocket sync services.
type SyncEngine struct {
	FirebaseClient  *firebase.FirebaseClient
	WebSocketServer *websocket.WebSocketServer
}

// NewSyncEngine initializes and returns a SyncEngine instance with Firebase and WebSocket services.
func NewSyncEngine(firebaseConfig *firebase.Config, wsAddr string) (*SyncEngine, error) {
	firebaseClient, err := firebase.NewClient(firebaseConfig)
	if err != nil {
		return nil, err
	}

	wsServer := websocket.NewServer()
	go wsServer.StartServer(wsAddr)

	return &SyncEngine{
		FirebaseClient:  firebaseClient,
		WebSocketServer: wsServer,
	}, nil
}

// SyncDataFromFirebase syncs data from Firebase to WebSocket clients.
func (se *SyncEngine) SyncDataFromFirebase() error {
	// Example Firebase data sync (customize for your use case)
	data, err := se.FirebaseClient.GetData("/some/path")
	if err != nil {
		return err
	}

	// Broadcast the data to WebSocket clients
	se.WebSocketServer.BroadcastMessage(data)
	return nil
}

// SyncDataToFirebase syncs data from WebSocket clients to Firebase.
func (se *SyncEngine) SyncDataToFirebase(client *websocket.WebSocketClient, data string) error {
	err := se.FirebaseClient.SetData("/some/path", data)
	if err != nil {
		return err
	}
	return nil
}

// StopSyncEngine shuts down the synchronization engine and its services.
func (se *SyncEngine) StopSyncEngine() {
	se.WebSocketServer.StopServer()
	se.FirebaseClient.Close()
}
