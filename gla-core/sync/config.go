package sync

// Config holds the configuration values for Firebase and WebSocket services.
type Config struct {
	FirebaseAPIKey  string
	FirebaseURL     string
	WebSocketAddr   string
}

// NewConfig creates a new configuration instance with the provided values.
func NewConfig(firebaseAPIKey, firebaseURL, webSocketAddr string) *Config {
	return &Config{
		FirebaseAPIKey: firebaseAPIKey,
		FirebaseURL:    firebaseURL,
		WebSocketAddr:  webSocketAddr,
	}
}
