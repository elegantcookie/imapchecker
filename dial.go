package imapchecker

import (
	"crypto/tls"
	"golang.org/x/net/context"
	"golang.org/x/net/proxy"
)

func DialWithContextDialer(ctx context.Context, dialer proxy.ContextDialer, checkerConfig *Config) (c *ImapConnection, err error) {
	err = parseConfig(checkerConfig)
	if err != nil {
		return
	}

	conn, err := dialer.DialContext(ctx, "tcp", checkerConfig.Address)
	if err != nil {
		return
	}

	if checkerConfig.UseTLS {
		conn = tls.Client(conn, checkerConfig.TLSConfig)
	}

	c = newConnection(conn, checkerConfig)
	err = c.verifyOnStart()
	return
}
