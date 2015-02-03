package mcpigo

import "net"

type Position struct {
	X, Y, Z float64
}

type MCPIConnection struct {
	World  World
	Player Player
	Chat   Chat
}

func Connect(host, port string) (MCPIConnection, error) {
	var mcConn MCPIConnection
	conn, err := net.Dial("tcp", host+":"+port)
	if err != nil {
		return MCPIConnection{}, err
	}

	mrw := NewMinecraftReadWriter(&MinecraftReader{conn}, &MinecraftWriter{conn})

	mcConn.World = NewWorld(*mrw)
	mcConn.Player = NewPlayer(*mrw)
	mcConn.Chat = NewChat(*mrw)

	return mcConn, err
}
