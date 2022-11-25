package imapchecker

import (
	"crypto/tls"
	"fmt"
	"golang.org/x/net/proxy"
	"net"
)

func DialWithDialer(dialer proxy.Dialer, domain string) (c *ImapConnection, err error) {
	address := domain + ":143"
	conn, err := dialer.Dial("tcp", address)
	if err != nil {
		return
	}
	c = NewImapConnection(conn, false, address)
	err = c.verifyOnStart()
	return
}

func DialWithDialerTLS(dialer proxy.Dialer, domain string, conf *tls.Config) (c *ImapConnection, err error) {
	address := domain + ":993"

	d, ok := dialer.(*net.Dialer)
	if !ok {
		err = fmt.Errorf("wrong dialer")
		return
	}

	conn, err := tls.DialWithDialer(d, "tcp", address, conf)
	if err != nil {
		return
	}
	c = NewImapConnection(conn, true, address)
	err = c.verifyOnStart()

	return c, err
}
