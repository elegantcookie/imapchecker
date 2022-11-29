package imapchecker

import (
	"net"
	"sync"
)

type ImapConnection struct {
	conn net.Conn

	m sync.Mutex

	state State
	conf  *Config
}

type State struct {
	Valid    bool
	LoggedIn bool
	Closed   bool
}
