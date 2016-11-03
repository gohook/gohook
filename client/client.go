package client

import (
	"fmt"
	"github.com/go-kit/kit/log"
	grpcclient "github.com/gohook/gohook-server/client"
	"github.com/gohook/gohook-server/gohookd"
	"github.com/gohook/gohook-server/pb"
	"github.com/gohook/gohook/config/configfile"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"io"
	"os"
	"os/exec"
	"strings"
	"time"
)

type Client interface {
	Tunnel() error
	Create()
	Remove()
	List()
}

type GohookClient struct {
	apiClient *grpcclient.GohookClient
	store     HookStore
}

func (c *GohookClient) Tunnel() error {
	stream, err := c.apiClient.Tunnel(context.Background(), &pb.TunnelRequest{})
	if err != nil {
		return nil
	}

	// Block and run forever
	for {
		response, err := stream.Recv()
		if err == io.EOF {
			fmt.Println("EOF")
			break
		}
		if err != nil {
			fmt.Println("Error", err)
			return err
		}
		hookCall := response.GetHook()

		// Load the latest from disk
		err = c.store.Load()
		if err != nil {
			fmt.Println("Error: ", err)
			continue
		}

		hook, err := c.store.Get(hookCall.Id)
		if err != nil {
			fmt.Println("Error: ", err)
			continue
		}

		fmt.Println("Running Hook: ", hook.Id)

		var rawCommand string
		if len(hook.Command) > 0 {
			rawCommand = strings.Join(hook.Command, " ")
		}

		cmd := exec.Command("sh", "-c", rawCommand)

		// TODO: Output to log file
		cmd.Stdout = os.Stdout
		cmd.Dir = hook.WorkingDir

		err = cmd.Run()
		if err != nil {
			fmt.Println("Failed to run command: ", rawCommand)
			continue
		}
	}
	return nil
}

func MergeHooks(localHooks HookList, remoteHooks gohookd.HookList) HookList {
	var hooks HookList
	localHookMap := make(map[string]*Hook)

	for _, h := range localHooks {
		localHookMap[h.Id] = h
	}

	for _, rh := range remoteHooks {
		if lh, ok := localHookMap[string(rh.Id)]; ok {
			// override some local fields with remote
			lh.Url = rh.Url
			hooks = append(hooks, lh)
		}
	}

	return hooks
}

func (c *GohookClient) List() (HookList, error) {
	var hooks HookList
	remoteHooks, err := c.apiClient.List(context.Background())
	if err != nil {
		return hooks, err
	}

	localHooks, err := c.store.List()
	if err != nil {
		return hooks, err
	}

	hooks = MergeHooks(localHooks, remoteHooks)
	return hooks, nil
}

func (c *GohookClient) Create(method string, command []string) (*Hook, error) {
	remoteHook, err := c.apiClient.Create(context.Background(), gohookd.HookRequest{method})
	if err != nil {
		return nil, err
	}

	workingDir, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	hook := &Hook{
		Id:         string(remoteHook.Id),
		Method:     remoteHook.Method,
		Url:        remoteHook.Url,
		Command:    command,
		WorkingDir: workingDir,
	}

	err = c.store.Add(hook)
	if err != nil {
		return nil, err
	}

	return hook, nil
}

func NewGohookClient(config *configfile.ConfigFile, store HookStore) (*GohookClient, error) {
	// Set up managed connection here
	conn, err := grpc.Dial(config.Host,
		grpc.WithInsecure(),
		grpc.WithBackoffMaxDelay(time.Second*10),
		grpc.WithPerRPCCredentials(&authToken{config.AuthToken}),
	)
	if err != nil {
		fmt.Println("error: ", err)
		return nil, err
	}

	service := grpcclient.New(conn, log.NewNopLogger())

	return &GohookClient{
		apiClient: &service,
		store:     store,
	}, nil
}
