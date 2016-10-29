package main

import (
	"golang.org/x/net/context"
)

type authToken struct {
	Token string
}

func (c *authToken) GetRequestMetadata(context.Context, ...string) (map[string]string, error) {
	return map[string]string{
		"token": c.Token,
	}, nil
}

func (c *authToken) RequireTransportSecurity() bool {
	return false
}
