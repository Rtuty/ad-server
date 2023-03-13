package main

import (
	"github.com/oschwald/geoip2-golang"
	"log"
	"modules/internal/ad-server"
)

func main() {
	geoip, err := geoip2.Open("GeoLite2-City.mmdb")
	if err != nil {
		log.Fatalf("open geoip db ERROR: %s", err)
	}

	s := adSrv.NewServer(geoip)
	if err := s.Listen(); err != nil {
		log.Fatalf("Server listening ERROR: %s", err)
	}

	log.Print("server are listening: \n https://localhost:8080")
}
