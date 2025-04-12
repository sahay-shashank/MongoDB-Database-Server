package handler

import (
	"net/http"

	"go.mongodb.org/mongo-driver/v2/mongo"
)

func WebHandler(w http.ResponseWriter, r *http.Request, client *mongo.Client) {
	w.Write([]byte("Hello World"))
}
