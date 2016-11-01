package filestore

import (
	"encoding/json"
	"errors"
	"github.com/gohook/gohook/client"
	"io"
	"os"
	"path/filepath"
	"sync"
)

type fileStore struct {
	Filename string
	mtx      sync.RWMutex
	hooks    client.HookList
}

func (f *fileStore) Add(hook *client.Hook) error {
	err := f.Load()
	if err != nil {
		return err
	}

	f.hooks = append(f.hooks, hook)

	return f.Write()
}

func (f *fileStore) List() (client.HookList, error) {
	err := f.Load()
	if err != nil {
		return client.HookList{}, err
	}
	return f.hooks, nil
}

func (f *fileStore) Get(id string) (*client.Hook, error) {
	err := f.Load()
	if err != nil {
		return nil, err
	}

	for _, h := range f.hooks {
		if h.Id == id {
			return h, nil
		}
	}

	return nil, errors.New("Not Found")
}

func (f *fileStore) LoadFromReader(hookStore io.Reader) error {
	return json.NewDecoder(hookStore).Decode(&f.hooks)
}

func (f *fileStore) Load() error {
	f.mtx.RLock()
	defer f.mtx.RUnlock()

	_, err := os.Stat(f.Filename)
	if err != nil {
		return err
	}
	file, err := os.Open(f.Filename)
	if err != nil {
		return err
	}
	defer file.Close()
	return f.LoadFromReader(file)
}

func (f *fileStore) SaveToWriter(writer io.Writer) error {
	data, err := json.Marshal(f.hooks)
	if err != nil {
		return err
	}
	_, err = writer.Write(data)
	return err
}

func (f *fileStore) Write() error {
	f.mtx.Lock()
	defer f.mtx.Unlock()

	// save the hooks to disk
	if err := os.MkdirAll(filepath.Dir(f.Filename), 0700); err != nil {
		return err
	}

	file, err := os.OpenFile(f.Filename, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		return err
	}

	defer file.Close()
	return f.SaveToWriter(file)
}

func NewLocalHookStore() client.HookStore {
	// Pass in opt data for settings and such for config
	return &fileStore{
		Filename: "/Users/begizi/.gohook/hooks.json",
	}
}