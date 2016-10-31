package client

import (
	"fmt"
	"github.com/go-kit/kit/log"
	grpcclient "github.com/gohook/gohook-server/client"
	"github.com/gohook/gohook-server/gohookd"
	"github.com/gohook/gohook-server/pb"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"time"
)

type Client interface {
	Create()
	Remove()
	List()
}

type GohookClient struct {
	apiClient *grpcclient.GohookClient
}

func (c *GohookClient) Tunnel() (pb.Gohook_TunnelClient, error) {
	return c.apiClient.Tunnel(context.Background(), &pb.TunnelRequest{})
}

func (c *GohookClient) List() (HookList, error) {
	var hooks HookList
	remoteHooks, err := c.apiClient.List(context.Background())
	if err != nil {
		return hooks, err
	}

	for _, h := range remoteHooks {
		hooks = append(hooks, &Hook{
			Id:     string(h.Id),
			Method: h.Method,
			Url:    h.Url,
		})
	}

	return hooks, nil
}

func (c *GohookClient) Create(method string, command []string) (*Hook, error) {
	remoteHook, err := c.apiClient.Create(context.Background(), gohookd.HookRequest{method})
	if err != nil {
		return nil, err
	}

	hook := &Hook{
		Id:     string(remoteHook.Id),
		Method: remoteHook.Method,
		Url:    remoteHook.Url,
	}

	return hook, nil
}

func NewGohookClient(token string) (*GohookClient, error) {
	// Set up managed connection here
	conn, err := grpc.Dial("localhost:9001",
		grpc.WithInsecure(),
		grpc.WithTimeout(time.Second),
		grpc.WithPerRPCCredentials(&authToken{token}),
	)
	if err != nil {
		fmt.Println("error: ", err)
		return nil, err
	}

	service := grpcclient.New(conn, log.NewNopLogger())

	return &GohookClient{
		apiClient: &service,
	}, nil
}
