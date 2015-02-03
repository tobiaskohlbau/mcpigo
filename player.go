package mcpigo

type Player interface {
	Position() Position
	PositionCoordinates() (int, int, int)
	SetPosition(pos Position)
	SetPositionFromCoordinates(x, y, z int)
}

type MCPIPlayer struct {
	rw MinecraftReadWriter
}

func NewPlayer(rw MinecraftReadWriter) Player {
	return MCPIPlayer{rw}
}

func (m MCPIPlayer) Position() Position {
	return m.rw.ReceivePosition()
}

func (m MCPIPlayer) PositionCoordinates() (int, int, int) {
	pos := m.Position()
	return int(pos.X), int(pos.Y), int(pos.Z)
}

func (m MCPIPlayer) SetPosition(pos Position) {
	m.rw.WritePosition(pos)
}

func (m MCPIPlayer) SetPositionFromCoordinates(x, y, z int) {
	m.rw.WritePosition(Position{float64(x), float64(y), float64(z)})
}
