package snmpfunc

import (
	"fmt"
	g "github.com/gosnmp/gosnmp"
	pb "snmpv2srv/proto"
	"sync"
	"time"
)

func BulkGet2Data(ip, community string, oids []string) (*pb.SnmpV2Result, error) {
	params := &g.GoSNMP{
		Target:    ip,
		Port:      uint16(161),
		Community: community,
		Version:   g.Version2c,
		Timeout:   time.Duration(2) * time.Second,
		//Logger:    g.NewLogger(log.New(os.Stdout, "", 0)),
	}
	var d *pb.SnmpV2Result = new(pb.SnmpV2Result)
	d.Ip = ip
	d.Community = community
	err := params.Connect()
	if err != nil {
		d.Date = time.Now().Format("2006年01月02日15:04:05")
		return d, err
	}
	defer params.Conn.Close()
	var data map[string]string = make(map[string]string)
	res, err := params.GetBulk(oids, 1, 255)
	for _, result := range res.Variables {
		switch result.Type {
		case g.OctetString:
			data[result.Name] = string(result.Value.([]byte))
		case g.ObjectIdentifier:
			continue
		default:
			data[result.Name] = fmt.Sprintf("%v", result.Value)
		}
	}
	d.Data = data
	d.Date = time.Now().Format("2006:01:02 15:04:05")
	return d, nil
}

/*
	并发获取设备信息，
*/
func SnmpV2BulkGet(rp *pb.SnmpV2BulkGetRequest) ([]*pb.SnmpV2Result, time.Duration) {
	startTime := time.Now()
	var wg sync.WaitGroup
	var res_ch chan *pb.SnmpV2Result = make(chan *pb.SnmpV2Result, len(rp.Ips))
	for _, ip := range rp.Ips {
		wg.Add(1)
		go func(ch chan<- *pb.SnmpV2Result, ipstr string, community string, oids []string) {
			defer wg.Done()
			result, _ := BulkGet2Data(ipstr, community, oids)
			ch <- result
		}(res_ch, ip, rp.Community, rp.Oids)
	}
	wg.Wait()
	close(res_ch)
	var ResponeData []*pb.SnmpV2Result = make([]*pb.SnmpV2Result, 0)
	for v := range res_ch {
		if v == nil {
			break
		}
		ResponeData = append(ResponeData, v)
	}
	endTime := time.Since(startTime)
	return ResponeData, endTime
}
