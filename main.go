package main

import (
	"fmt"
	"github.com/gerschkin/config"
	"github.com/gerschkin/server/rpc"
	"github.com/valyala/gorpc"
	"gopkg.in/alecthomas/kingpin.v2"
	"os"
)

var (
	conf   config.Client
	client *gorpc.Client
)

func main() {
	// Parse flags
	kingpin.Parse()

	// Try and read the configuration file. If we can't read the file then
	// we exit with exit code 1.
	err := config.ReadClient(*configPath, &conf)
	if err != nil {
		fmt.Println("Could not read config file:", err)
		os.Exit(1)
	}

	// Register rpc types that we accept, and figure out the address to use
	// with the rpc server.
	registerTypes()
	addr := fmt.Sprintf("%s:%d", conf.RPC.Host, conf.RPC.Port)

	// Create our TCP client and start it.
	client = gorpc.NewTCPClient(addr)
	client.Start()

	defer client.Stop()

	// Send a test
	fmt.Println(send(rpc.Request{
		Type: rpc.REQUEST_PING,
	}))
}

// We need to register all of the types that our rpc server can
// send and receive. Normal types such as int, string, ... are
// automatically registered and work.
func registerTypes() {
	gorpc.RegisterType(rpc.Request{})
	gorpc.RegisterType(rpc.Response{})
}

func send(request rpc.Request) (interface{}, error) {
	return client.Call(request)
}
