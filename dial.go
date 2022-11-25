package imapchecker

import (
	"crypto/tls"
	"golang.org/x/net/proxy"
)

func DialWithDialer(dialer proxy.Dialer, address string) (c *ImapConnection, err error) {
	conn, err := dialer.Dial("tcp", address)
	if err != nil {
		return
	}
	c = NewImapConnection(conn, false, address)
	err = c.verifyOnStart()
	return
}

func DialWithDialerTLS(dialer proxy.Dialer, address string, conf *tls.Config) (c *ImapConnection, err error) {

	d, err := dialer.Dial("tcp", address)
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
