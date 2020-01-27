package sink

import (
	"fmt"
	"sync"
)

var (
	sMu sync.Mutex
	s   = make(map[string]func(Config) (Sink, error))
)

func Register(name string, setupFunc func(Config) (Sink, error)) {
	sMu.Lock()
	defer sMu.Unlock()
	if _, dup := s[name]; dup {
		panic("sink: Register called twice for sink " + name)
	}
	s[name] = setupFunc
}

func New(c Config) (Sink, error) {
	setupFunc, ok := s[c.Kind]
	if !ok {
		return nil, fmt.Errorf("store: store '%s' does not exist", c.Kind)
	}
	return setupFunc(c)
}

type Sink interface {
	Send(string) error
}

type Config struct {
	Kind       string `yaml:"kind"`
	Connection string `yaml:"connection"`
}
