package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"scalingmq-go/cmd/broker/app"
	"scalingmq-go/common/http/native"
	"scalingmq-go/core/broker/httphandler"
	"time"
)

var commandLine *flag.FlagSet

// init 启动初始化加载
func init() {
	fmt.Println("scaling mq broker init.")

	commandLine = flag.NewFlagSet("broker", flag.ExitOnError)
	commandLine.Usage = func() {
		fmt.Println("this is scaling mq broker usage")
		commandLine.PrintDefaults()
	}

	p := app.GetInstance()
	p.Parse(commandLine)
}

func main() {
	rand.Seed(time.Now().UnixNano())
	_ = commandLine.Parse(os.Args[1:])
	// 打印加载的参数
	p := app.GetInstance()
	fmt.Printf("系统启动,加载的参数=%+v\n", *p)

	if *p.HttpNative {
		server := native.HttpServer{}
		err := server.StartServer(func() {
			httphandler.HandleHttpReq()
		})
		if err != nil {
			return
		}
	} else {

	}
}
