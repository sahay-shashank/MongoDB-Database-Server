package handler

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/v2/mongo"
)

func CloseDatabase(client *mongo.Client) error {
	if err := client.Disconnect(context.TODO()); err != nil {
		return fmt.Errorf("error disconnecting from MongoDB. Error: %v", err)
	}
	return nil
}
