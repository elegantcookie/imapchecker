package imapchecker

import (
	"fmt"
	"net"
	"strings"
	"sync"
	"time"
)

func (c *ImapConnection) GetState() State {
	return c.state
}

func (c *ImapConnection) CheckInbox(login, password string) (err error) {
	err = c.Authenticate(login, password)
	if err != nil {
		return
	}
	err = c.OpenInbox()
	if err != nil {
		return
	}
	return
}

func (c *ImapConnection) OpenInbox() (err error) {
	res, err := c.WriteMessage(`tag select "INBOX"`)
	if err != nil {
		return
	}

	if strings.Index(res, "OK") == -1 {
		err = PermissionDenied
		return
	}

	c.m.Lock()
	c.state.Valid = true
	c.m.Unlock()

	return
}

func (c *ImapConnection) Authenticate(login, password string) (err error) {
	res, err := c.WriteMessage("authenticate login " + login + " " + password)
	if err != nil {
		return
	}

	if strings.Index(res, "authenticate OK") == -1 {
		err = WrongCredentials
		return
	}

	c.m.Lock()
	c.state.LoggedIn = true
	c.m.Unlock()

	return
}

func (c *ImapConnection) WriteMessage(message string) (res string, err error) {
	if c.Closed() {
		err = ConnectionClosed
		return
	}

	_, err = c.conn.Write([]byte(message + "\r\n"))
	if err != nil {
		return
	}

	return c.ReadMessage()

}

func (c *ImapConnection) ReadMessage() (res string, err error) {
	if c.Closed() {
		err = ConnectionClosed
		return
	}
	buf := make([]byte, 512)
	_, err = c.conn.Read(buf)

	if err != nil {
		return
	}

	res = string(buf)
	return

}

func (c *ImapConnection) verifyOnStart() (err error) {
	res, err := c.ReadMessage()
	if err != nil {
		return
	}
	if strings.Index(res, "OK") == -1 {
		err = fmt.Errorf("failed to connect the server")
	}
	return
}

func NewImapConnection(conn net.Conn, tls bool, address string) *ImapConnection {
	return &ImapConnection{
		conn:    conn,
		tls:     tls,
		address: address,
		state:   State{},
	}
}

func (c *ImapConnection) Closed() bool {
	return c.state.Closed
}

func (c *ImapConnection) Read(b []byte) (n int, err error) {
	return c.conn.Read(b)
}

func (c *ImapConnection) Write(b []byte) (n int, err error) {
	return c.conn.Write(b)
}

func (c *ImapConnection) Close() error {
	var m sync.Mutex
	m.Lock()
	c.state.Closed = true
	m.Unlock()
	return c.conn.Close()
}

func (c *ImapConnection) LocalAddr() net.Addr {
	return c.conn.LocalAddr()
}

func (c *ImapConnection) RemoteAddr() net.Addr {
	return c.conn.RemoteAddr()
}

func (c *ImapConnection) SetDeadline(t time.Time) error {
	return c.conn.SetDeadline(t)
}

func (c *ImapConnection) SetReadDeadline(t time.Time) error {
	return c.conn.SetReadDeadline(t)
}

func (c *ImapConnection) SetWriteDeadline(t time.Time) error {
	return c.conn.SetWriteDeadline(t)
}
