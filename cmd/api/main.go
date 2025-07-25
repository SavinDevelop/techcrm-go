package main

import (
	"context"
	"github.com/SavinDevelop/techcrm-go/internal/transport"
	"github.com/SavinDevelop/techcrm-go/pkg/db"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	pg, err := db.NewPostgres()
	if err != nil {
		log.Fatalf("DB init error: %v", err)
	}

	server := transport.NewHTTPServer(pg)
	go server.Start()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Failed to shutdown server: %v", err)
	}
}
