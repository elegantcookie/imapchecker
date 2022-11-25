package imapchecker

import "errors"

var (
	ConnectionClosed = errors.New("connection is closeed")
	WrongCredentials = errors.New("wrong credentials")
	PermissionDenied = errors.New("permission denied")
)
