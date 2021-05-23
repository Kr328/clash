package constant

import "net"

type WrappedConn interface {
	RawConn() net.Conn
}

type WrappedPacketConn interface {
	RawPacketConn() net.PacketConn
}