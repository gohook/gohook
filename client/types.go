package client

type HookList []*Hook

type Hook struct {
	Id         string
	Method     string
	Url        string
	Command    []string
	WorkingDir string
	OutDir     string
}
