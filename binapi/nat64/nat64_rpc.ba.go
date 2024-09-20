// Code generated by GoVPP's binapi-generator. DO NOT EDIT.

package nat64

import (
	"context"
	"fmt"
	"io"

	api "github.com/alkiranet/govpp/api"
	memclnt "github.com/alkiranet/govpp/binapi/memclnt"
)

// RPCService defines RPC service nat64.
type RPCService interface {
	Nat64AddDelInterface(ctx context.Context, in *Nat64AddDelInterface) (*Nat64AddDelInterfaceReply, error)
	Nat64AddDelInterfaceAddr(ctx context.Context, in *Nat64AddDelInterfaceAddr) (*Nat64AddDelInterfaceAddrReply, error)
	Nat64AddDelPoolAddrRange(ctx context.Context, in *Nat64AddDelPoolAddrRange) (*Nat64AddDelPoolAddrRangeReply, error)
	Nat64AddDelPrefix(ctx context.Context, in *Nat64AddDelPrefix) (*Nat64AddDelPrefixReply, error)
	Nat64AddDelStaticBib(ctx context.Context, in *Nat64AddDelStaticBib) (*Nat64AddDelStaticBibReply, error)
	Nat64BibDump(ctx context.Context, in *Nat64BibDump) (RPCService_Nat64BibDumpClient, error)
	Nat64GetTimeouts(ctx context.Context, in *Nat64GetTimeouts) (*Nat64GetTimeoutsReply, error)
	Nat64InterfaceDump(ctx context.Context, in *Nat64InterfaceDump) (RPCService_Nat64InterfaceDumpClient, error)
	Nat64PluginEnableDisable(ctx context.Context, in *Nat64PluginEnableDisable) (*Nat64PluginEnableDisableReply, error)
	Nat64PoolAddrDump(ctx context.Context, in *Nat64PoolAddrDump) (RPCService_Nat64PoolAddrDumpClient, error)
	Nat64PrefixDump(ctx context.Context, in *Nat64PrefixDump) (RPCService_Nat64PrefixDumpClient, error)
	Nat64SetTimeouts(ctx context.Context, in *Nat64SetTimeouts) (*Nat64SetTimeoutsReply, error)
	Nat64StDump(ctx context.Context, in *Nat64StDump) (RPCService_Nat64StDumpClient, error)
}

type serviceClient struct {
	conn api.Connection
}

func NewServiceClient(conn api.Connection) RPCService {
	return &serviceClient{conn}
}

func (c *serviceClient) Nat64AddDelInterface(ctx context.Context, in *Nat64AddDelInterface) (*Nat64AddDelInterfaceReply, error) {
	out := new(Nat64AddDelInterfaceReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) Nat64AddDelInterfaceAddr(ctx context.Context, in *Nat64AddDelInterfaceAddr) (*Nat64AddDelInterfaceAddrReply, error) {
	out := new(Nat64AddDelInterfaceAddrReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) Nat64AddDelPoolAddrRange(ctx context.Context, in *Nat64AddDelPoolAddrRange) (*Nat64AddDelPoolAddrRangeReply, error) {
	out := new(Nat64AddDelPoolAddrRangeReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) Nat64AddDelPrefix(ctx context.Context, in *Nat64AddDelPrefix) (*Nat64AddDelPrefixReply, error) {
	out := new(Nat64AddDelPrefixReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) Nat64AddDelStaticBib(ctx context.Context, in *Nat64AddDelStaticBib) (*Nat64AddDelStaticBibReply, error) {
	out := new(Nat64AddDelStaticBibReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) Nat64BibDump(ctx context.Context, in *Nat64BibDump) (RPCService_Nat64BibDumpClient, error) {
	stream, err := c.conn.NewStream(ctx)
	if err != nil {
		return nil, err
	}
	x := &serviceClient_Nat64BibDumpClient{stream}
	if err := x.Stream.SendMsg(in); err != nil {
		return nil, err
	}
	if err = x.Stream.SendMsg(&memclnt.ControlPing{}); err != nil {
		return nil, err
	}
	return x, nil
}

type RPCService_Nat64BibDumpClient interface {
	Recv() (*Nat64BibDetails, error)
	api.Stream
}

type serviceClient_Nat64BibDumpClient struct {
	api.Stream
}

func (c *serviceClient_Nat64BibDumpClient) Recv() (*Nat64BibDetails, error) {
	msg, err := c.Stream.RecvMsg()
	if err != nil {
		return nil, err
	}
	switch m := msg.(type) {
	case *Nat64BibDetails:
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

func (c *serviceClient) Nat64GetTimeouts(ctx context.Context, in *Nat64GetTimeouts) (*Nat64GetTimeoutsReply, error) {
	out := new(Nat64GetTimeoutsReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) Nat64InterfaceDump(ctx context.Context, in *Nat64InterfaceDump) (RPCService_Nat64InterfaceDumpClient, error) {
	stream, err := c.conn.NewStream(ctx)
	if err != nil {
		return nil, err
	}
	x := &serviceClient_Nat64InterfaceDumpClient{stream}
	if err := x.Stream.SendMsg(in); err != nil {
		return nil, err
	}
	if err = x.Stream.SendMsg(&memclnt.ControlPing{}); err != nil {
		return nil, err
	}
	return x, nil
}

type RPCService_Nat64InterfaceDumpClient interface {
	Recv() (*Nat64InterfaceDetails, error)
	api.Stream
}

type serviceClient_Nat64InterfaceDumpClient struct {
	api.Stream
}

func (c *serviceClient_Nat64InterfaceDumpClient) Recv() (*Nat64InterfaceDetails, error) {
	msg, err := c.Stream.RecvMsg()
	if err != nil {
		return nil, err
	}
	switch m := msg.(type) {
	case *Nat64InterfaceDetails:
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

func (c *serviceClient) Nat64PluginEnableDisable(ctx context.Context, in *Nat64PluginEnableDisable) (*Nat64PluginEnableDisableReply, error) {
	out := new(Nat64PluginEnableDisableReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) Nat64PoolAddrDump(ctx context.Context, in *Nat64PoolAddrDump) (RPCService_Nat64PoolAddrDumpClient, error) {
	stream, err := c.conn.NewStream(ctx)
	if err != nil {
		return nil, err
	}
	x := &serviceClient_Nat64PoolAddrDumpClient{stream}
	if err := x.Stream.SendMsg(in); err != nil {
		return nil, err
	}
	if err = x.Stream.SendMsg(&memclnt.ControlPing{}); err != nil {
		return nil, err
	}
	return x, nil
}

type RPCService_Nat64PoolAddrDumpClient interface {
	Recv() (*Nat64PoolAddrDetails, error)
	api.Stream
}

type serviceClient_Nat64PoolAddrDumpClient struct {
	api.Stream
}

func (c *serviceClient_Nat64PoolAddrDumpClient) Recv() (*Nat64PoolAddrDetails, error) {
	msg, err := c.Stream.RecvMsg()
	if err != nil {
		return nil, err
	}
	switch m := msg.(type) {
	case *Nat64PoolAddrDetails:
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

func (c *serviceClient) Nat64PrefixDump(ctx context.Context, in *Nat64PrefixDump) (RPCService_Nat64PrefixDumpClient, error) {
	stream, err := c.conn.NewStream(ctx)
	if err != nil {
		return nil, err
	}
	x := &serviceClient_Nat64PrefixDumpClient{stream}
	if err := x.Stream.SendMsg(in); err != nil {
		return nil, err
	}
	if err = x.Stream.SendMsg(&memclnt.ControlPing{}); err != nil {
		return nil, err
	}
	return x, nil
}

type RPCService_Nat64PrefixDumpClient interface {
	Recv() (*Nat64PrefixDetails, error)
	api.Stream
}

type serviceClient_Nat64PrefixDumpClient struct {
	api.Stream
}

func (c *serviceClient_Nat64PrefixDumpClient) Recv() (*Nat64PrefixDetails, error) {
	msg, err := c.Stream.RecvMsg()
	if err != nil {
		return nil, err
	}
	switch m := msg.(type) {
	case *Nat64PrefixDetails:
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

func (c *serviceClient) Nat64SetTimeouts(ctx context.Context, in *Nat64SetTimeouts) (*Nat64SetTimeoutsReply, error) {
	out := new(Nat64SetTimeoutsReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) Nat64StDump(ctx context.Context, in *Nat64StDump) (RPCService_Nat64StDumpClient, error) {
	stream, err := c.conn.NewStream(ctx)
	if err != nil {
		return nil, err
	}
	x := &serviceClient_Nat64StDumpClient{stream}
	if err := x.Stream.SendMsg(in); err != nil {
		return nil, err
	}
	if err = x.Stream.SendMsg(&memclnt.ControlPing{}); err != nil {
		return nil, err
	}
	return x, nil
}

type RPCService_Nat64StDumpClient interface {
	Recv() (*Nat64StDetails, error)
	api.Stream
}

type serviceClient_Nat64StDumpClient struct {
	api.Stream
}

func (c *serviceClient_Nat64StDumpClient) Recv() (*Nat64StDetails, error) {
	msg, err := c.Stream.RecvMsg()
	if err != nil {
		return nil, err
	}
	switch m := msg.(type) {
	case *Nat64StDetails:
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
