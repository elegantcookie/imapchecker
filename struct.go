package imapchecker

import (
	"net"
	"sync"
)

type ImapConnection struct {
	conn net.Conn

	tls     bool
	address string

	state State

	m sync.Mutex
}

type State struct {
	LoggedIn bool
	Closed   bool
}
