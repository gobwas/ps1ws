package jsonrpc

import (
	"bufio"
	"net"
)

// DECL> OMIT
// Channel represents some absctract channel over net.Conn.
type Channel struct {
	conn net.Conn
	out  chan Packet
}

func NewChannel(conn net.Conn) *Channel {
	c := &Channel{conn, make(chan Packet, N)}
	go c.reader()
	go c.writer()
}

// DECL< OMIT
// IMPL> OMIT
func (c *Channel) reader() {
	buf := bufio.NewReader(c.conn) // Allocation.
	for {
		readPacket(buf) // Possibly allocations too.
		// ...
	}
}

func (c *Channel) writer() {
	buf := bufio.NewWriter(c.conn) // Allocation.
	for pkt := range c.out {
		writePacket(buf, pkt)
		// ...
		buf.Flush()
	}
}

// IMPL< OMIT
