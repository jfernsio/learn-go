package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/jfernsio/learn-go/internal/config"
)

func main() {
	//loa config
	cfg := config.MustLoad()
	//db setup
	//setup router
	router := http.NewServeMux()

	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome to api"))
	})
	//setup server
	server := http.Server{
		Addr : cfg.Addr,
		Handler : router,
	}

	
	fmt.Printf("Starting server on %s\n", cfg.Addr)

	done := make(chan os.Signal,1)

	signal.Notify(done, os.Interrupt,syscall.SIGINT, syscall.SIGTERM)

	go func () {
		err := server.ListenAndServe()

		if err != nil {
			log.Fatal("failed to start server")
		}
	} ()
	
	//wait for shutdown signal
	<-done

	slog.Printf("Shutting down server on %s\n", cfg.Addr)
}