package main

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

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

	slog.Info("Starting server",slog.String("address",cfg.Addr))

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

	slog.Info("Shutting down server")

	ctx , cancel := context.WithTimeout(context.Background(),5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		slog.Error("failed to shutdown server",slog.String("error",err.Error()))
	}

	slog.Info("server shutdown gracefully")
}