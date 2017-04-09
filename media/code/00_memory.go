package jsonrpc

import (
	"bufio"
	"net"
)

// START OMIT
// Channel represents some absctract channel over net.Conn.
type Channel struct {
	conn net.Conn
	out  chan Packet
}

func NewChannel(conn net.Conn) *Channel {
	c := &Channel{conn}
	go c.reader() // Stack growth.
	go c.writer() // Stack growth.
}

func (c *Channel) reader() {
	buf := bufio.NewReader(c.conn) // Allocation.
	for {
		readPacket(buf) // Possibly allocations too.
		// ...
	}
}

// END OMIT

func (c *Channel) writer() {
	buf := bufio.NewWriter(c.conn) // Allocation.
	for pkt := range c.out {
		writePacket(buf, pkt)
		// ...
		buf.Flush()
	}
}
