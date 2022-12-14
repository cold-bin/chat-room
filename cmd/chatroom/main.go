package main

import (
	"fmt"
	"log"
	"net/http"

	_ "net/http/pprof"

	"github.com/polaris1119/chatroom/global"
	"github.com/polaris1119/chatroom/server"
)

var (
	addr   = ":8088"
	banner = `
    ____              _____
   |    |    |   /\     |
   |    |____|  /  \    | 
   |    |    | /----\   |
   |____|    |/      \  |		ChatRoom，start on：%s
`
)

func init() {
	global.Init()
}

func main() {
	fmt.Printf(banner, addr)

	server.RegisterHandle()

	log.Fatal(http.ListenAndServe(addr, nil))
}
