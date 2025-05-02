package firebase

import (
	"context"
	"fmt"

	"firebase.google.com/go/v4"
    "firebase.google.com/go/v4/auth"
    "firebase.google.com/go/v4/db"
)

type AuthClient struct {
	client *auth.Client
}

func NewAuthClient(app *firebase.App) (*AuthClient, error) {
	client, err := app.Auth(context.Background())
	if err != nil {
		return nil, fmt.Errorf("error getting Firebase Auth client: %v", err)
	}
	return &AuthClient{client: client}, nil
}

func (ac *AuthClient) CreateUser(email, password string) (*auth.UserRecord, error) {
	userParams := (&auth.UserToCreate{}).Email(email).Password(password)
	userRecord, err := ac.client.CreateUser(context.Background(), userParams)
	if err != nil {
		return nil, fmt.Errorf("error creating user: %v", err)
	}
	return userRecord, nil
}

func (ac *AuthClient) VerifyIDToken(idToken string) (*auth.Token, error) {
	token, err := ac.client.VerifyIDToken(context.Background(), idToken)
	if err != nil {
		return nil, fmt.Errorf("error verifying ID token: %v", err)
	}
	return token, nil
}
