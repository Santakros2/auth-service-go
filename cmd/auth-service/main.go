package main

import (
	"auth-service/internal/config"
	"auth-service/internal/db"
	"auth-service/internal/handlers"
	"auth-service/internal/repository"
	"auth-service/internal/service"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	config := config.LoadConfigMySQL()
	db, err := db.ConnecMysql(config)

	if err != nil {
		log.Fatal("databse error ", err)
		return
	}
	defer db.Close()

	repo := repository.NewUserRepository(db)

	service := service.NewService(repo)

	handler := handlers.NewHandler(service)

	mux := http.NewServeMux()

	mux.HandleFunc("/login", handler.LoginHandle)
	mux.HandleFunc("/logout", handler.Logout)

	server := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	go func() {
		log.Println("Server running on port 8080.")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen error %v", err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	<-stop

	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	server.Shutdown(ctx)

}
