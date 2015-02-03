package mcpigo

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"strings"
)

type Connection interface {
	Send(s string) (n int, err error)
	Receive() (r string, err error)
	World() (World, error)
	Player() (Player, error)
}

type Position struct {
	X, Y, Z int
}

type PositionFloat struct {
	X, Y, Z float64
}

type MCPIConnection struct {
	conn io.ReadWriter
}

func (m MCPIConnection) Send(s string) (n int, err error) {
	n, err = fmt.Fprintf(m.conn, "%s\n", s)
	return
}

func (m MCPIConnection) Receive() (r string, err error) {
	r, err = bufio.NewReader(m.conn).ReadString('\n')
	if err == nil {
		r = strings.TrimSuffix(r, "\n")
	}
	return
}

func (m MCPIConnection) World() (World, error) {
	world := new(MCPIWorld)
	world.conn = m
	return world, nil
}

func (m MCPIConnection) Player() (Player, error) {
	player := new(MCPIPlayer)
	player.conn = m
	return player, nil
}

func Connect(host, port string) (Connection, error) {
	var piConn MCPIConnection
	conn, err := net.Dial("tcp", host+":"+port)
	if err != nil {
		return nil, err
	}
	piConn.conn = conn
	return piConn, err
}
