package main

import (
	"log"

	"github.com/imishinist/ayb/pkg/bot"
)

func main() {
	s := bot.CreateServer()
	log.Fatal(s.ListenAndServe(":8090"))
}
