package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"students-api/internal/config"
	"students-api/internal/handlers/student"
	"syscall"
	"time"
)

func main() {
	// load config

	cfg := config.MustLoad()

	// database setup

	// setup router
	router := http.NewServeMux()

	router.HandleFunc("POST /api/students", student.New())

	// setup server

	server := http.Server{
		Addr:    cfg.Address,
		Handler: router,
	}

	fmt.Println("server started", cfg.HTTPServer.Address)

	done := make(chan os.Signal, 1)

	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// 	Why a goroutine?

	// Because server.ListenAndServe() is blocking.
	go func() {
		// graceful shutdown
		err := server.ListenAndServe()

		if err != nil {
			log.Fatal("failed to start server")
		}
	}()

	<-done // blocking

	// server stop logic

	slog.Info("shutting down the server")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := server.Shutdown(ctx)

	if err != nil {
		slog.Error("failed to shutdown the server", slog.String("error", err.Error()))
	}

	slog.Info("server shut down succesfully")

}
