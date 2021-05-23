package dialer

import (
	"context"
	"net"
)

type TunnelDialer func(context context.Context, network, address string) (net.Conn, error)

var DefaultTunnelDialer TunnelDialer

func DialTunnelContext(ctx context.Context, network, address string) (net.Conn, error) {
	if dialer := DefaultTunnelDialer; dialer != nil {
		return dialer(ctx, network, address)
	}

	return DialContext(ctx, network, address)
}
