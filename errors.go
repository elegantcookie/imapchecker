package imapchecker

import "errors"

var (
	ConnectionClosed = errors.New("connection is closeed")
	NotAuthenticated = errors.New("not authenticated")
)
