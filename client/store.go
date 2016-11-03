package client

type HookStore interface {
	Add(hook *Hook) error
	List() (HookList, error)
	Get(id string) (*Hook, error)
	Remove(id string) error
	Write() error
	Load() error
}
