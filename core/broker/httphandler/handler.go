package httphandler

import (
	"fmt"
	"log"
	"net/http"
)

type Handle func()

func HandleHttpReq() {
	http.HandleFunc("/test", func(writer http.ResponseWriter, request *http.Request) {
		log.Println("收到请求...")
		if request.Body != http.NoBody {
			readBytes := make([]byte, 128)
			read, err := request.Body.Read(readBytes)
			if err != nil {
				fmt.Println(err)
				return
			}
			if read > 0 {
				// 解析http body内容
			}
		}
		id := request.FormValue("id")
		fmt.Printf("id:%v\n", id)
		_, err := writer.Write([]byte("ok"))
		if err != nil {
			fmt.Println(err)
			return
		}
	})
}
