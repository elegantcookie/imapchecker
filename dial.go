package imapchecker

import (
	"crypto/tls"
	"golang.org/x/net/context"
	"golang.org/x/net/proxy"
)

func DialWithContextDialer(ctx context.Context, dialer proxy.ContextDialer, address string) (c *ImapConnection, err error) {
	conn, err := dialer.DialContext(ctx, "tcp", address)
	if err != nil {
		return
	}
	c = NewImapConnection(conn, false, address)
	err = c.verifyOnStart()
	return
}

func DialWithContextDialerTLS(ctx context.Context, dialer proxy.ContextDialer, address string, conf *tls.Config) (c *ImapConnection, err error) {

	d, err := dialer.DialContext(ctx, "tcp", address)
	if err != nil {
		return
	}
	conn := tls.Client(d, conf)
	if err != nil {
		return
	}
	c = NewImapConnection(conn, true, address)
	err = c.verifyOnStart()

	return c, err
}
