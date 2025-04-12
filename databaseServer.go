package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/mainframematrix/MongoDB-Database-Server/database"
	handler "github.com/mainframematrix/MongoDB-Database-Server/handler"
)

func main() {
	client, err := database.ConnectDatabase()
	if err != nil {
		log.Fatalf("MongoDB initialization failed: %v", err)
		panic(err)
	}
	defer database.CloseDatabase(client)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-quit
		log.Println("Server shutting down...")
		database.CloseDatabase(client)
		os.Exit(0)
	}()

	server := http.NewServeMux()
	server.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		handler.WebHandler(w, r, client)
	})
	http.ListenAndServe(":8081", server)
}
