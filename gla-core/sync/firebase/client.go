package firebase

import (
	"context"
	"fmt"
	"log"

	"firebase.google.com/go/v4"
	"google.golang.org/api/option"
	"github.com/firebase/firebase-admin-go/"
	
)

type FirebaseClient struct {
	app *firebase.App
}

func NewClient(serviceAccountFile string) (*FirebaseClient, error) {
	opt := option.WithCredentialsFile(serviceAccountFile)
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return nil, fmt.Errorf("error initializing app: %v", err)
	}
	return &FirebaseClient{app: app}, nil
}

func (fc *FirebaseClient) GetFirestoreClient() (*firebase.FirestoreClient, error) {
	client, err := fc.app.Firestore(context.Background())
	if err != nil {
		return nil, fmt.Errorf("error getting Firestore client: %v", err)
	}
	return client, nil
}
