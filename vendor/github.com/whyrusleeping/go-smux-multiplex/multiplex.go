package peerstream_multiplex

import (
	"net"

	mp "github.com/libp2p/go-mplex"          // Conn is a connection to a remote peer.
	smux "github.com/libp2p/go-stream-muxer" // Conn is a connection to a remote peer.
	"fmt"
)

type conn struct {
	*mp.Multiplex
}

func (c *conn) Close() error {
	return c.Multiplex.Close()
}

func (c *conn) IsClosed() bool {
	return c.Multiplex.IsClosed()
}

// OpenStream creates a new stream.
func (c *conn) OpenStream() (smux.Stream, error) {
	return c.Multiplex.NewStream()
}

// AcceptStream accepts a stream opened by the other side.
func (c *conn) AcceptStream() (smux.Stream, error) {
	return c.Multiplex.Accept()
}

// Transport is a go-peerstream transport that constructs
// multiplex-backed connections.
type Transport struct{}

// DefaultTransport has default settings for multiplex
var DefaultTransport = &Transport{}

func (t *Transport) NewConn(nc net.Conn, isServer bool) (smux.Conn, error) {
	fmt.Printf("#### multiplex:Transport.NewConn Called \n")
	return &conn{mp.NewMultiplex(nc, isServer)}, nil
}
