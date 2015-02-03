package mcpigo

type Chat struct {
	rw MinecraftReadWriter
}

func NewChat(rw MinecraftReadWriter) Chat {
	return Chat{rw}
}

func (c Chat) Message(s string) {
	c.rw.WriteMessage(s)
}
