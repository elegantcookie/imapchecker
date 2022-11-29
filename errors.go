package imapchecker

import "errors"

var (
	ConnectionClosed = errors.New("connection is closed")
	WrongCredentials = errors.New("wrong credentials")
	PermissionDenied = errors.New("permission denied")
	EmptyAddress     = errors.New("empty address")
	NoConfig         = errors.New("no config provided")
	NoTLSConfig      = errors.New("no tls config provided with useTLS flag")
)
