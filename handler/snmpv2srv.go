package handler

import (
	"context"
	"fmt"
	pb "snmpv2srv/proto"
	snmp "snmpv2srv/snmpfunc"
)

type Snmpv2srv struct{}

//子叶点Get
func (e *Snmpv2srv) SnmpV2Get(ctx context.Context, req *pb.SnmpV2GetRequest, rsp *pb.SnmpV2Response) error {
	res, date := snmp.SnmpV2Get(req)
	rsp.Result = res
	rsp.Date = fmt.Sprintf("%v", date)
	return nil
}

//非子叶点Get
func (e *Snmpv2srv) SnmpV2BulkGet(ctx context.Context, req *pb.SnmpV2BulkGetRequest, rsp *pb.SnmpV2Response) error {
	res, date := snmp.SnmpV2BulkGet(req)
	rsp.Result = res
	rsp.Date = fmt.Sprintf("%v", date)
	return nil
}
