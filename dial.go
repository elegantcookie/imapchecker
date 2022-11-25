package imapchecker

import (
	"crypto/tls"
	"net"
)

func DialWithDialer(dialer *net.Dialer, domain string) (c *ImapConnection, err error) {
	address := domain + ":143"
	conn, err := dialer.Dial("tcp", address)
	if err != nil {
		return
	}
	c = NewImapConnection(conn, false, address)
	err = c.verifyOnStart()
	return
}

func DialWithDialerTLS(dialer *net.Dialer, domain string, conf *tls.Config) (c *ImapConnection, err error) {
	address := domain + ":993"
	conn, err := tls.DialWithDialer(dialer, "tcp", address, conf)
	if err != nil {
		return
	}
	c = NewImapConnection(conn, true, address)
	err = c.verifyOnStart()

	return c, err
}
