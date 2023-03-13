package main

import (
	"log"
	"modules/internal/ad-server"
)

func main() {
	s := adSrv.NewServer()
	if err := s.Listen(); err != nil {
		log.Fatalf("Server listening ERROR: %s", err)
	}
}
