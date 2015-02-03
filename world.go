package mcpigo

type World interface {
	BlockAtPosition(pos Position) int
	BlockAtCoordinates(x, y, z int) int
	SetBlockAtPosition(pos Position, id int)
	SetBlockAtCoordinates(x, y, z, id int)
	SetBlocksInPositionRange(startPos, endPos Position, id int)
	SetBlocksInCoordinatesRange(xs, ys, zs, xe, ye, ze, id int)
}

type MCPIWorld struct {
	rw MinecraftReadWriter
}

func NewWorld(rw MinecraftReadWriter) World {
	return MCPIWorld{rw}
}

func (m MCPIWorld) BlockAtPosition(pos Position) int {
	return m.rw.ReceiveBlock(pos)
}

func (m MCPIWorld) BlockAtCoordinates(x, y, z int) int {
	return m.BlockAtPosition(Position{float64(x), float64(y), float64(z)})
}

func (m MCPIWorld) SetBlockAtPosition(pos Position, id int) {
	m.rw.WriteBlock(pos, id)
}

func (m MCPIWorld) SetBlockAtCoordinates(x, y, z, id int) {
	m.SetBlockAtPosition(Position{float64(x), float64(y), float64(z)}, id)
}

func (m MCPIWorld) SetBlocksInPositionRange(startPos, endPos Position, id int) {
	m.rw.WriteBlocks(startPos, endPos, id)
}

func (m MCPIWorld) SetBlocksInCoordinatesRange(xs, ys, zs, xe, ye, ze, id int) {
	m.SetBlocksInPositionRange(Position{float64(xs), float64(ys), float64(zs)}, Position{float64(xe), float64(ye), float64(ze)}, id)
}
