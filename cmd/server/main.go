package main

import (
	"github.com/oceanho/gw"
	"github.com/oceanho/gw/conf"
	"github.com/oceanho/gwpro/app"
)

func main() {
	bcs := conf.NewBootConfigFromFile("config/boot.yaml")
	opts := gw.NewServerOption(bcs)
	server := gw.NewServerWithOption(opts)
	registerApps(server)
	server.Serve()
}

func registerApps(server *gw.HostServer) {
	server.Register(app.New())
}
