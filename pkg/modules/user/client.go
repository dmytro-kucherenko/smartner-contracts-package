package user

import (
	"context"

	"github.com/dmytro-kucherenko/smartner-contracts-package/gen/go/user"
	adapter "github.com/dmytro-kucherenko/smartner-utils-package/pkg/server/adapters/grpc"
	"google.golang.org/grpc"
)

type Client struct {
	client user.UserClient
}

func NewClient(conn *grpc.ClientConn) *Client {
	return &Client{user.NewUserClient(conn)}
}

func (c *Client) Get(ctx context.Context, params GetParams) (result GetResult, err error) {
	return adapter.HandleCall(c.client.Get, ctx, params, result)
}
