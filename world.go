package mcpigo

import (
	"fmt"
	"strconv"
)

type World interface {
	BlockAtCoordinates(x, y, z int) int
	BlockAtPosition(pos Position) int
}

type MCPIWorld struct {
	conn Connection
}

func (m MCPIWorld) BlockAtPosition(pos Position) int {
	return m.BlockAtCoordinates(pos.X, pos.Y, pos.Z)
}

func (m MCPIWorld) BlockAtCoordinates(x, y, z int) int {
	m.conn.Send(fmt.Sprintf("world.getBlock(%v,%v,%v)", x, y, z))
	resp, err := m.conn.Receive()
	if err != nil {
		return -1
	}
	id, err := strconv.Atoi(resp)
	if err != nil {
		return -1
	}
	return id
}
