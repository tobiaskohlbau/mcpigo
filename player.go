package mcpigo

import (
	"fmt"
	"strconv"
	"strings"
)

type Player interface {
	Position() Position
	PositionFloat() PositionFloat
	SetPositionFromCoordinates(x, y, z int)
	SetPosition(pos Position)
}

type MCPIPlayer struct {
	conn Connection
}

func (m MCPIPlayer) Position() Position {
	posF := m.PositionFloat()
	return Position{int(posF.X), int(posF.Y), int(posF.Z)}
}

func (m MCPIPlayer) PositionFloat() PositionFloat {
	m.conn.Send("player.getPos()")
	resp, err := m.conn.Receive()
	if err != nil {
		return PositionFloat{}
	}
	posList := strings.Split(resp, ",")
	var pos [3]float64
	for i, s := range posList {
		tmp, _ := strconv.ParseFloat(s, 64)
		pos[i] = tmp
	}
	return PositionFloat{pos[0], pos[1], pos[2]}
}

func (m MCPIPlayer) SetPositionFromCoordinates(x, y, z int) {
	m.conn.Send(fmt.Sprintf("player.setPos(%v,%v,%v)", x, y, z))
}

func (m MCPIPlayer) SetPosition(pos Position) {
	m.SetPositionFromCoordinates(pos.X, pos.Y, pos.Z)
}
