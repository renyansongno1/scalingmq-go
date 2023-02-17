package native

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"scalingmq-go/cmd/broker/app"
	"scalingmq-go/core/broker/httphandler"
)

type HttpServer struct {
	srv *http.Server
}

func (HttpServer *HttpServer) StartServer(handle httphandler.Handle) error {
	// 注册handler
	handle()

	p := app.GetInstance()
	server := &http.Server{Addr: fmt.Sprintf(":%v", *p.HttpPort), Handler: nil}
	HttpServer.srv = server
	go log.Fatal(server.ListenAndServe())
	return nil
}

func (HttpServer *HttpServer) StopServer() error {
	return HttpServer.srv.Shutdown(context.Background())
}
