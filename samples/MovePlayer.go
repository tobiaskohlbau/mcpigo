package main

import (
	"fmt"
	"log"
	"os"

	"github.com/tobiaskohlbau/mcpigo"
)

func main() {
	mc, err := mcpigo.Connect("192.168.10.164", "4711")
	if err != nil {
		log.Fatal(err)
		os.Exit(-1)
	}

	mc.Chat.Message("Getting positon...")
	pos := mc.Player.Position()
	x, y, z := mc.Player.PositionCoordinates()
	pos.X += 5
	mc.Chat.Message("Changing position...")
	mc.Player.SetPosition(pos)

	pos.X -= 5
	id := mc.World.BlockAtPosition(pos)
	fmt.Println(id)

	mc.World.SetBlockAtCoordinates(x, y+5, z, 1)

	pos1 := pos
	pos2 := pos
	pos2.X += 5

	mc.World.SetBlocksInPositionRange(pos1, pos2, 1)
}
