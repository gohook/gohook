package flags

type ClientOptions struct {
	Debug    bool
	LogLevel string
	Version  bool
}

func NewClientOptions() *ClientOptions {
	return &ClientOptions{}
}
