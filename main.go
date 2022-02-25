package main

import (
	"github.com/asim/go-micro/plugins/registry/consul/v4"
	"go-micro.dev/v4"
	log "go-micro.dev/v4/logger"
	"snmpv2srv/handler"
	pb "snmpv2srv/proto"

	"go-micro.dev/v4/registry"
)

var (
	service = "snmpv2"
	version = "latest"
)

func main() {
	consulRegistry := consul.NewRegistry(
		func(op *registry.Options) {
			op.Addrs = []string{
				"172.28.102.216:8500",
			}
		},
	)
	// Create service
	srv := micro.NewService(
		micro.Registry(consulRegistry),
		micro.Name(service),
		micro.Version(version),
		micro.Address("172.28.102.216:9009"),
	)
	srv.Init()

	// Register handler
	pb.RegisterSnmpv2SrvHandler(srv.Server(), new(handler.Snmpv2srv))
	// Run service
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}
