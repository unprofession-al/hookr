package decoder

import (
	"fmt"
	"sync"
)

var (
	dMu sync.Mutex
	d   = make(map[string]func(Config) (Decoder, error))
)

func Register(name string, setupFunc func(Config) (Decoder, error)) {
	dMu.Lock()
	defer dMu.Unlock()
	if _, dup := d[name]; dup {
		panic("decoder: Register called twice for docoder " + name)
	}
	d[name] = setupFunc
}

func New(c Config) (Decoder, error) {
	setupFunc, ok := d[c.Kind]
	if !ok {
		return nil, fmt.Errorf("decoder: Decoder '%s' does not exist", c.Kind)
	}
	return setupFunc(c)
}

type Decoder interface {
	Decode([]byte) ([]byte, error)
}

type Config struct {
	Kind string `yaml:"kind"`
}
