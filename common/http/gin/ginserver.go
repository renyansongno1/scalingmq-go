package gin

import (
	"context"
	"fmt"
	"net/http"
	"scalingmq-go/cmd/broker/app"
	"scalingmq-go/core/broker/httphandler"
)
import "github.com/gin-gonic/gin"

type Server struct {
	Engine *gin.Engine
	srv    *http.Server
}

func (Server *Server) StartServer(handle httphandler.Handle) error {
	r := gin.Default()
	Server.Engine = r
	p := app.GetInstance()
	server := &http.Server{Addr: fmt.Sprintf(":%v", *p.HttpPort), Handler: nil}
	Server.srv = server
	go func() {
		err := r.Run(fmt.Sprintf(":%v", *p.HttpPort))
		if err != nil {

		}
	}()

	return nil
}

func (Server *Server) StopServer() error {
	return Server.srv.Shutdown(context.Background())
}
