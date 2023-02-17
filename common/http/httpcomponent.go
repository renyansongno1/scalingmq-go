package http

import "scalingmq-go/core/broker/httphandler"

type Server interface {
	StartServer(handle httphandler.Handle) error
	StopServer() error
}
