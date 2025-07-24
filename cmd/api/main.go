package main

import (
	"github.com/SavinDevelop/techcrm-go/pkg/db"
	"log"
)

func main() {
	pg, err := db.NewPostgres()
	if err != nil {
		log.Fatalf("DB init error: %v", err)
	}
	defer func() {
		if err := pg.Close(); err != nil {
			log.Printf("Error closing DB: %v", err)
		}
	}()
}
