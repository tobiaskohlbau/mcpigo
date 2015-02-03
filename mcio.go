package mcpigo

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"strconv"
	"strings"
)

type MinecraftReader struct {
	r io.Reader
}

func NewMinecraftReader(r io.Reader) MinecraftReader {
	return MinecraftReader{r}
}

func (mc MinecraftReader) ReadPosition() Position {
	resp, err := mc.Receive()
	if err != nil {
		fmt.Println(err)
		return Position{}
	}
	posList := strings.Split(resp, ",")
	var pos [3]float64
	for i, s := range posList {
		tmp, _ := strconv.ParseFloat(s, 64)
		pos[i] = tmp
	}
	return Position{pos[0], pos[1], pos[2]}
}

func (mc MinecraftReader) ReadBlock() int {
	resp, err := mc.Receive()
	if err != nil {
		fmt.Println(err)
		return -1
	}
	fmt.Println(resp)
	id, err := strconv.Atoi(resp)
	if err != nil {
		log.Fatal(err)
		return -1
	}
	return id
}

func (mc MinecraftReader) Receive() (r string, err error) {
	r, err = bufio.NewReader(mc.r).ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return
	}
	r = strings.TrimSuffix(r, "\n")
	return
}

type MinecraftWriter struct {
	w io.Writer
}

func (mw MinecraftWriter) WritePosition(pos Position) {
	s := fmt.Sprintf("player.setPos(%v,%v,%v", pos.X, pos.Y, pos.Z)
	mw.Send(s)
}

func (mw MinecraftWriter) WriteMessage(message string) {
	s := fmt.Sprintf("chat.post(%s)", message)
	mw.Send(s)
}

func (mw MinecraftWriter) WriteBlock(position Position, id int) {
	s := fmt.Sprintf("world.setBlock(%v,%v,%v,%v)", int(position.X), int(position.Y), int(position.Z), id)
	mw.Send(s)
}

func (mw MinecraftWriter) WriteBlocks(startPos, endPos Position, id int) {
	s := fmt.Sprintf("world.setBlocks(%v,%v,%v,%v,%v,%v,%v)", int(startPos.X), int(startPos.Y), int(startPos.Z), int(endPos.X), int(endPos.Y), int(endPos.Z), id)
	mw.Send(s)
}

func (mw MinecraftWriter) Send(s string) (n int, err error) {
	n, err = fmt.Fprintf(mw.w, "%s\n", s)
	return
}

type MinecraftReadWriter struct {
	*MinecraftReader
	*MinecraftWriter
}

func NewMinecraftReadWriter(r *MinecraftReader, w *MinecraftWriter) *MinecraftReadWriter {
	return &MinecraftReadWriter{r, w}
}

func (mrw MinecraftReadWriter) ReceivePosition() Position {
	_, err := mrw.Send("player.getPos()")
	if err != nil {
		fmt.Println(err)
	}
	return mrw.ReadPosition()
}

func (mrw MinecraftReadWriter) ReceiveBlock(position Position) int {
	s := fmt.Sprintf("world.getBlock(%v,%v,%v)", int(position.X), int(position.Y), int(position.Z))
	_, err := mrw.Send(s)
	if err != nil {
		fmt.Println(err)
		return -1
	}
	return mrw.ReadBlock()
}
