package main

import (
	"log"
	"os"

	"github.com/tobiaskohlbau/mcpigo"
)

func main() {
	conn, err := mcpigo.Connect("192.168.10.164", "4711")
	if err != nil {
		log.Fatal(err)
		os.Exit(-1)
	}

	//world, err := conn.World()
	//if err != nil {
	//	log.Fatal(err)
	//	os.Exit(-1)
	//}

	player, err := conn.Player()
	if err != nil {
		log.Fatal(err)
		os.Exit(-1)
	}

	position := player.Position()
	position.X += 5
	player.SetPosition(position)
}
