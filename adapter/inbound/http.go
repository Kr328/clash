package inbound

import (
	"net"

	C "github.com/Dreamacro/clash/constant"
	"github.com/Dreamacro/clash/context"
	"github.com/Dreamacro/clash/transport/socks5"
)

// NewHTTP receive normal http request and return HTTPContext
func NewHTTP(target string, rawSrc, rawDst net.Addr, conn net.Conn) *context.ConnContext {
	metadata := parseSocksAddr(socks5.ParseAddr(target))
	metadata.NetWork = C.TCP
	metadata.Type = C.HTTP
	if ip, port, err := parseAddr(rawSrc.String()); err == nil {
		metadata.SrcIP = ip
		metadata.SrcPort = port
	}

	metadata.RawSrcAddr = rawSrc
	metadata.RawDstAddr = rawDst

	return context.NewConnContext(conn, metadata)
}