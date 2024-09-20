// Code generated by GoVPP's binapi-generator. DO NOT EDIT.

package dslite

import (
	"context"
	"fmt"
	"io"

	api "github.com/alkiranet/govpp/api"
	memclnt "github.com/alkiranet/govpp/binapi/memclnt"
)

// RPCService defines RPC service dslite.
type RPCService interface {
	DsliteAddDelPoolAddrRange(ctx context.Context, in *DsliteAddDelPoolAddrRange) (*DsliteAddDelPoolAddrRangeReply, error)
	DsliteAddressDump(ctx context.Context, in *DsliteAddressDump) (RPCService_DsliteAddressDumpClient, error)
	DsliteGetAftrAddr(ctx context.Context, in *DsliteGetAftrAddr) (*DsliteGetAftrAddrReply, error)
	DsliteGetB4Addr(ctx context.Context, in *DsliteGetB4Addr) (*DsliteGetB4AddrReply, error)
	DsliteSetAftrAddr(ctx context.Context, in *DsliteSetAftrAddr) (*DsliteSetAftrAddrReply, error)
	DsliteSetB4Addr(ctx context.Context, in *DsliteSetB4Addr) (*DsliteSetB4AddrReply, error)
}

type serviceClient struct {
	conn api.Connection
}

func NewServiceClient(conn api.Connection) RPCService {
	return &serviceClient{conn}
}

func (c *serviceClient) DsliteAddDelPoolAddrRange(ctx context.Context, in *DsliteAddDelPoolAddrRange) (*DsliteAddDelPoolAddrRangeReply, error) {
	out := new(DsliteAddDelPoolAddrRangeReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) DsliteAddressDump(ctx context.Context, in *DsliteAddressDump) (RPCService_DsliteAddressDumpClient, error) {
	stream, err := c.conn.NewStream(ctx)
	if err != nil {
		return nil, err
	}
	x := &serviceClient_DsliteAddressDumpClient{stream}
	if err := x.Stream.SendMsg(in); err != nil {
		return nil, err
	}
	if err = x.Stream.SendMsg(&memclnt.ControlPing{}); err != nil {
		return nil, err
	}
	return x, nil
}

type RPCService_DsliteAddressDumpClient interface {
	Recv() (*DsliteAddressDetails, error)
	api.Stream
}

type serviceClient_DsliteAddressDumpClient struct {
	api.Stream
}

func (c *serviceClient_DsliteAddressDumpClient) Recv() (*DsliteAddressDetails, error) {
	msg, err := c.Stream.RecvMsg()
	if err != nil {
		return nil, err
	}
	switch m := msg.(type) {
	case *DsliteAddressDetails:
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

func (c *serviceClient) DsliteGetAftrAddr(ctx context.Context, in *DsliteGetAftrAddr) (*DsliteGetAftrAddrReply, error) {
	out := new(DsliteGetAftrAddrReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) DsliteGetB4Addr(ctx context.Context, in *DsliteGetB4Addr) (*DsliteGetB4AddrReply, error) {
	out := new(DsliteGetB4AddrReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) DsliteSetAftrAddr(ctx context.Context, in *DsliteSetAftrAddr) (*DsliteSetAftrAddrReply, error) {
	out := new(DsliteSetAftrAddrReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) DsliteSetB4Addr(ctx context.Context, in *DsliteSetB4Addr) (*DsliteSetB4AddrReply, error) {
	out := new(DsliteSetB4AddrReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}
