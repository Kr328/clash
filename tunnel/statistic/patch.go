package statistic

import (
	"net"
)

func (m *Manager) Total() (up, down int64) {
	return m.uploadTotal.Load(), m.downloadTotal.Load()
}

func (tt *tcpTracker) RawConn() net.Conn {
	if tt.Chain.Last() == "DIRECT" {
		return tt.Conn
	}

	return nil
}

func (ut *udpTracker) RawPacketConn() net.PacketConn {
	if ut.Chain.Last() == "DIRECT" {
		return ut.PacketConn
	}

	return nil
}
