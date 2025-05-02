package main

import (
	"fmt"
	"log"
	"gla-core/security"
	"gla-core/sync"
	"gla-core/sandbox"
	"gla-core/llm"
	"gla-core/api"

	firebaseSync "sync/firebase" // Replace with your package path
)

func main() {
	// Initialize the security manager
	securityManager, err := security.NewSecurityManager("audit.log")
	if err != nil {
		log.Fatalf("Error initializing security manager: %v", err)
	}
	defer securityManager.Close()

	// Initialize synchronization services
	firebaseSync := sync.NewFirebaseSync()
	websocketSync := sync.NewWebSocketSync()

	// Start sandbox execution engines (Yaegi, WASM, Docker)
	yaegiSandbox := sandbox.NewYaegiSandbox()
	wasmSandbox := sandbox.NewWasmSandbox()
	dockerSandbox := sandbox.NewDockerSandbox()

	// Example of using LLM for code generation
	LLMModel, err := llm.NewLLMModel("codellama-go.bin")
	if err != nil {
		log.Fatalf("Error loading LLM model: %v", err)
	}

	// Generate some Go code with the LLM
	code, err := LLMModel.GenerateCode("Create a function that returns the square of a number")
	if err != nil {
		log.Fatalf("Error generating code: %v", err)
	}

	// Log and sanitize generated code
	sanitizedCode := securityManager.HandleSecurity(code, "Generated code from LLM model")
	fmt.Println("Sanitized Generated Code:", sanitizedCode)

	// Start the WebSocket server
	go websocketSync.StartServer("localhost:8080")

	// Simulate interaction with Firebase
	err = firebaseSync.Connect("firebase-project-id")
	if err != nil {
		log.Fatalf("Error connecting to Firebase: %v", err)
	}

	// Initialize Firebase
	err := firebaseSync.Connect("sync/gla-core-firebase-adminsdk-fbsvc-f7247032a0.json") // Path to your JSON file
  if err != nil {
    log.Fatalf("Failed to initialize Firebase: %v", err)
  }

	// Print status to show the components are initialized
	fmt.Println("System initialized successfully.")
}
