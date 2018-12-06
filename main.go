package main

import (
	"fmt"
	"github.com/wiggle-penguin/wp-mmo-be-core/io"
	"github.com/wiggle-penguin/wp-mmo-be-core/mem"
)

var (
	ListenAddress = ":8080"
)

func main() {

	fmt.Println("✅  Starting memory database")
	mem.Init()

	fmt.Println("⁉️  TODO: Starting etcd integration")
	//TODO: etcd client init

	fmt.Println("✅  Starting websocket server")
	io.StartWebSocketServer(ListenAddress)

}
