package tunnel

import (
	"net"

	"github.com/Dreamacro/clash/constant"
)

func unwrap(conn net.Conn) net.Conn {
	r := conn

	for r != nil {
		if c, ok := r.(constant.WrappedConn); ok {
			r = c.RawConn()
		} else {
			break
		}
	}

	if r == nil {
		return conn
	}

	return r
}

func unwrapPacket(conn net.PacketConn) net.PacketConn {
	r := conn

	for r != nil {
		if c, ok := r.(constant.WrappedPacketConn); ok {
			r = c.RawPacketConn()
		} else {
			break
		}
	}

	if r == nil {
		return conn
	}

	return r
}