README.md - Firebase-Specific Setup and Usage
markdown
Copy
Edit
# Firebase Integration for GLA Core

This directory contains the necessary functionality to integrate Firebase services into the GLA Core engine. It includes the Firebase client setup, authentication, and real-time database interactions.

## Setup

1. Install the Firebase Go SDK:

```bash
go get firebase.google.com/go
go get google.golang.org/api/option
Ensure you have the service account key JSON file from Firebase and use it to initialize the Firebase app.

Firebase Client
NewClient(serviceAccountFile string) initializes the Firebase app with the provided service account JSON file. You can use this to interact with Firebase services like Firestore and Realtime Database.

Firebase Authentication
NewAuthClient(app *firebase.App) creates a new authentication client. You can use it to create users, verify ID tokens, and perform other authentication tasks.

Create User:
go
Copy
Edit
authClient, err := firebase.NewAuthClient(firebaseApp)
if err != nil {
	log.Fatalf("Failed to initialize Auth client: %v", err)
}
user, err := authClient.CreateUser("user@example.com", "securepassword")
if err != nil {
	log.Fatalf("Failed to create user: %v", err)
}
fmt.Println("Created user:", user.UID)
Firebase Realtime Database
NewDatabaseClient(app *firebase.App) creates a new database client for interacting with Firebase Realtime Database. You can set and get data using this client.

Set Data:
go
Copy
Edit
dbClient, err := firebase.NewDatabaseClient(firebaseApp)
if err != nil {
	log.Fatalf("Failed to initialize Database client: %v", err)
}
data := map[string]string{"message": "Hello, Firebase!"}
err = dbClient.SetData("/greetings", data)
if err != nil {
	log.Fatalf("Failed to set data: %v", err)
}
Get Data:
go
Copy
Edit
var result map[string]string
err = dbClient.GetData("/greetings", &result)
if err != nil {
	log.Fatalf("Failed to get data: %v", err)
}
fmt.Println("Retrieved data:", result)