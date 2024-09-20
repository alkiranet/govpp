// Code generated by GoVPP's binapi-generator. DO NOT EDIT.

package gre

import (
	"context"
	"fmt"
	"io"

	api "github.com/alkiranet/govpp/api"
	memclnt "github.com/alkiranet/govpp/binapi/memclnt"
)

// RPCService defines RPC service gre.
type RPCService interface {
	GreTunnelAddDel(ctx context.Context, in *GreTunnelAddDel) (*GreTunnelAddDelReply, error)
	GreTunnelDump(ctx context.Context, in *GreTunnelDump) (RPCService_GreTunnelDumpClient, error)
}

type serviceClient struct {
	conn api.Connection
}

func NewServiceClient(conn api.Connection) RPCService {
	return &serviceClient{conn}
}

func (c *serviceClient) GreTunnelAddDel(ctx context.Context, in *GreTunnelAddDel) (*GreTunnelAddDelReply, error) {
	out := new(GreTunnelAddDelReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) GreTunnelDump(ctx context.Context, in *GreTunnelDump) (RPCService_GreTunnelDumpClient, error) {
	stream, err := c.conn.NewStream(ctx)
	if err != nil {
		return nil, err
	}
	x := &serviceClient_GreTunnelDumpClient{stream}
	if err := x.Stream.SendMsg(in); err != nil {
		return nil, err
	}
	if err = x.Stream.SendMsg(&memclnt.ControlPing{}); err != nil {
		return nil, err
	}
	return x, nil
}

type RPCService_GreTunnelDumpClient interface {
	Recv() (*GreTunnelDetails, error)
	api.Stream
}

type serviceClient_GreTunnelDumpClient struct {
	api.Stream
}

func (c *serviceClient_GreTunnelDumpClient) Recv() (*GreTunnelDetails, error) {
	msg, err := c.Stream.RecvMsg()
	if err != nil {
		return nil, err
	}
	switch m := msg.(type) {
	case *GreTunnelDetails:
		return m, nil
	case *memclnt.ControlPingReply:
		err = c.Stream.Close()
		if err != nil {
			return nil, err
		}
		return nil, io.EOF
	default:
		return nil, fmt.Errorf("unexpected message: %T %v", m, m)
	}
}
