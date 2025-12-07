package main

import (
	"fmt"
	"log"
	"net/http"
	"students-api/internal/config"
)

func main() {
	// load config

	cfg := config.MustLoad()

	// database setup

	// setup router
	router := http.NewServeMux()

	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome to students api"))
	})

	// setup server

	server := http.Server{
		Addr:    cfg.Address,
		Handler: router,
	}

	fmt.Println("server started" , cfg.HTTPServer.Address)

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("failed to start server")
	}
}
