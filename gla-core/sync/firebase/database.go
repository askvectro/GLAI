package firebase

import (
	"context"
	"fmt"

	"firebase.google.com/go/v4"
    "firebase.google.com/go/v4/auth"
    "firebase.google.com/go/v4/db"
)

type DatabaseClient struct {
	client *db.Client
}

func NewDatabaseClient(app *firebase.App) (*DatabaseClient, error) {
	client, err := app.Database(context.Background())
	if err != nil {
		return nil, fmt.Errorf("error getting Firebase Database client: %v", err)
	}
	return &DatabaseClient{client: client}, nil
}

func (dc *DatabaseClient) SetData(path string, data interface{}) error {
	ref := dc.client.NewRef(path)
	err := ref.Set(context.Background(), data)
	if err != nil {
		return fmt.Errorf("error setting data in Firebase: %v", err)
	}
	return nil
}

func (dc *DatabaseClient) GetData(path string, result interface{}) error {
	ref := dc.client.NewRef(path)
	err := ref.Get(context.Background(), &result)
	if err != nil {
		return fmt.Errorf("error getting data from Firebase: %v", err)
	}
	return nil
}
