package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/asim/go-micro/plugins/registry/consul/v4"
	"go-micro.dev/v4"
	"go-micro.dev/v4/client"
	"go-micro.dev/v4/registry"
	pb "snmpv2srv/proto"
	"time"
)

func main() {
	t1 := time.Now()
	consulRegistry := consul.NewRegistry(
		func(op *registry.Options) {
			op.Addrs = []string{
				"172.28.102.216:8500",
			}
		},
	)
	service := micro.NewService(
		micro.Registry(consulRegistry),
	)
	service.Init()
	elapsed1 := time.Since(t1)
	fmt.Println("App elapsed: ", elapsed1)
	clientService := pb.NewSnmpv2SrvService("snmpv2", service.Client())
	var opss client.CallOption = func(o *client.CallOptions) {
		o.RequestTimeout = time.Second * 30
		o.DialTimeout = time.Second * 30
	}
	res, err := clientService.SnmpV2Get(context.TODO(), &pb.SnmpV2GetRequest{
		Ips:       []string{"192.168.115.128"},
		Community: "elovedotaer",
		Data:      map[string]string{"memTotalSwap": ".1.3.6.1.4.1.2021.4.3.0", "SysDescGET": ".1.3.6.1.2.1.1.1.0", "sysUptimeGET": ".1.3.6.1.2.1.1.3.0", "sysContactGET": ".1.3.6.1.2.1.1.4.0", "SysNameGET": ".1.3.6.1.2.1.1.5.0", "SysLocationGET": ".1.3.6.1.2.1.1.6.0", "SysServiceGET": ".1.3.6.1.2.1.1.7.0", "hrSWRunNameWALK": ".1.3.6.1.2.1.25.4.2.1.2", "hrSWInstalledNameWALK": ".1.3.6.1.2.1.25.6.3.1.2"},
	}, opss)
	if err != nil {
		fmt.Println("err:", err)
	}
	t, _ := json.Marshal(res)
	fmt.Printf("res:%#v\n", string(t))
	elapsed2 := time.Since(t1)
	fmt.Println("App elapsed: ", elapsed2)
}
