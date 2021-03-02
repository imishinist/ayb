package main

import (
	"fmt"
	"log"
	"os"

	"github.com/imishinist/ayb/pkg/bot"
)

func main() {
	s := bot.CreateServer()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8090"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	log.Fatal(s.ListenAndServe(fmt.Sprintf(":%s", port)))
}
