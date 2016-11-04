package flags

import (
	"github.com/gohook/gohook/config"
)

type ClientOptions struct {
	Debug    bool
	LogLevel string
	Version  bool
	HookDir  string
}

var DefaultClientOptions = ClientOptions{
	LogLevel: "warn",
	HookDir:  config.ConfigDir(),
}

func NewClientOptions() *ClientOptions {
	return &DefaultClientOptions
}
