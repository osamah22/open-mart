package main

import (
	"log"
	"strconv"

	_ "github.com/lib/pq"
)

func main() {
	// 1️⃣ Load config
	cfg, err := LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	// 2️⃣ Create server (this wires EVERYTHING)
	srv, err := NewServer(cfg)
	if err != nil {
		log.Fatal(err)
	}

	// 3️⃣ Run server
	if err := srv.Router.RunTLS("localhost:"+strconv.Itoa(cfg.PORT), "./cert/server.crt", "./cert/server.key"); err != nil {
		log.Fatal(err)
	}
}
