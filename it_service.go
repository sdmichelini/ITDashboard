package main

import (
	"time"
	//"os"
	//"bytes"
	"fmt"
	//"net"
	//"golang.org/x/net/icmp"
	//"golang.org/x/net/ipv4"
)

//Time out for Ping
var PingTimeout = 3000 * time.Millisecond

type ItService struct{
	Name string //Name of the Service
	IpAddress string //IpAddress of the Service
	Id uint //What type the service is
	Status bool //Is the Serivce Up?
}


//Inspired by https://golang.org/src/net/ipraw_test.go
/*
@function PingService
@param service
	ItService to ping
@return success
	True if the Item is Up, False if not
@return error
	nil if no error, else the error
*/
func PingService(service ItService) (bool, error){
	/*
	c, err := net.Dial("ip4:icmp", service.IpAddress)
	//Make sure we can Dial the remote
	if err != nil {
		return false, err
	}
	//Set the Deadline
	c.SetDeadline(time.Now().Add(PingTimeout))

	//Close the socket when we are done
	defer c.Close()

	xid, xseq := os.Getpid()&0xffff,1

	wb,err := (icmpMessage{
		Type: icmpv4EchoRequest, Code: 0,
		Body: &icmpEcho{
			ID:xid, Seq: xseq,
			Data: bytes.Repeat([]byte("IT Ping"), 3),
		},
		}).Marshal()

	if err != nil {
		return false,err
	}

	if _,err := c.Write(wb); err != nil{
		return false,err
	}

	var m *ipv4.icmpMessage
	rb := make([]byte, 20+len(wb))
	for {
		if _, err := c.Read(rb); err != nil {
			return false,err
		}
		if net == "ip4" {
			rb = ipv4Payload(rb)
		}	
		if m, err = parseICMPMessage(rb); err != nil {
			return false, err
		}
		switch m.Type {
		case ipv4.icmpv4EchoRequest, ipv4.icmpv6EchoRequest:
			continue
		}
		break
	}
	switch p := m.Body.(type) {
	case *ipv4.icmpEcho:
		if p.ID != xid || p.Seq != xseq {
			return false, fmt.Errorf("got id=%v, seqnum=%v; expected id=%v, seqnum=%v", p.ID, p.Seq, xid, xseq)
		}
		return true, nil
	default:
		return false, fmt.Errorf("got type=%v, code=%v; expected type=%v, code=%v", m.Type, m.Code, typ, 0)
	}
	*/
	return false, fmt.Errorf("Not Implemented")
}