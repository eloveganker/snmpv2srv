// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/snmpv2srv.proto

package proto

import (
	fmt "fmt"
	proto "google.golang.org/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "go-micro.dev/v4/api"
	client "go-micro.dev/v4/client"
	server "go-micro.dev/v4/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for Snmpv2Srv service

func NewSnmpv2SrvEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Snmpv2Srv service

type Snmpv2SrvService interface {
	SnmpV2Get(ctx context.Context, in *SnmpV2GetRequest, opts ...client.CallOption) (*SnmpV2Response, error)
	SnmpV2BulkGet(ctx context.Context, in *SnmpV2BulkGetRequest, opts ...client.CallOption) (*SnmpV2Response, error)
}

type snmpv2SrvService struct {
	c    client.Client
	name string
}

func NewSnmpv2SrvService(name string, c client.Client) Snmpv2SrvService {
	return &snmpv2SrvService{
		c:    c,
		name: name,
	}
}

func (c *snmpv2SrvService) SnmpV2Get(ctx context.Context, in *SnmpV2GetRequest, opts ...client.CallOption) (*SnmpV2Response, error) {
	req := c.c.NewRequest(c.name, "Snmpv2Srv.SnmpV2Get", in)
	out := new(SnmpV2Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *snmpv2SrvService) SnmpV2BulkGet(ctx context.Context, in *SnmpV2BulkGetRequest, opts ...client.CallOption) (*SnmpV2Response, error) {
	req := c.c.NewRequest(c.name, "Snmpv2Srv.SnmpV2BulkGet", in)
	out := new(SnmpV2Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Snmpv2Srv service

type Snmpv2SrvHandler interface {
	SnmpV2Get(context.Context, *SnmpV2GetRequest, *SnmpV2Response) error
	SnmpV2BulkGet(context.Context, *SnmpV2BulkGetRequest, *SnmpV2Response) error
}

func RegisterSnmpv2SrvHandler(s server.Server, hdlr Snmpv2SrvHandler, opts ...server.HandlerOption) error {
	type snmpv2Srv interface {
		SnmpV2Get(ctx context.Context, in *SnmpV2GetRequest, out *SnmpV2Response) error
		SnmpV2BulkGet(ctx context.Context, in *SnmpV2BulkGetRequest, out *SnmpV2Response) error
	}
	type Snmpv2Srv struct {
		snmpv2Srv
	}
	h := &snmpv2SrvHandler{hdlr}
	return s.Handle(s.NewHandler(&Snmpv2Srv{h}, opts...))
}

type snmpv2SrvHandler struct {
	Snmpv2SrvHandler
}

func (h *snmpv2SrvHandler) SnmpV2Get(ctx context.Context, in *SnmpV2GetRequest, out *SnmpV2Response) error {
	return h.Snmpv2SrvHandler.SnmpV2Get(ctx, in, out)
}

func (h *snmpv2SrvHandler) SnmpV2BulkGet(ctx context.Context, in *SnmpV2BulkGetRequest, out *SnmpV2Response) error {
	return h.Snmpv2SrvHandler.SnmpV2BulkGet(ctx, in, out)
}
