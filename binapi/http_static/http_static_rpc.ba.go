// Code generated by GoVPP's binapi-generator. DO NOT EDIT.

package http_static

import (
	"context"

	api "github.com/alkiranet/govpp/api"
)

// RPCService defines RPC service http_static.
type RPCService interface {
	HTTPStaticEnable(ctx context.Context, in *HTTPStaticEnable) (*HTTPStaticEnableReply, error)
}

type serviceClient struct {
	conn api.Connection
}

func NewServiceClient(conn api.Connection) RPCService {
	return &serviceClient{conn}
}

func (c *serviceClient) HTTPStaticEnable(ctx context.Context, in *HTTPStaticEnable) (*HTTPStaticEnableReply, error) {
	out := new(HTTPStaticEnableReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}
